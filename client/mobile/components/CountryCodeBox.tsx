import React, { useEffect, useState } from "react";
import {
  KeyboardAvoidingView,
  Modal,
  Platform,
  Pressable,
  ScrollView,
  StyleSheet,
  Text,
  TextInput,
  View,
} from "react-native";
import { Controller, useFormContext } from "react-hook-form";
import countries from "i18n-iso-countries";
import enLocale from "i18n-iso-countries/langs/en.json";
import {
  CountryCode as LibCountryCode,
  getCountryCallingCode,
} from "libphonenumber-js";
import { colors } from "@/constants";

countries.registerLocale(enLocale);

type FormValue = {
  countryCode: LibCountryCode;
  phoneNumber: string;
};

type CountryItem = {
  cca2: LibCountryCode;
  name: string;
  callingCode: string;
  label: string;
};

function buildCountryItems(): CountryItem[] {
  const names = countries.getNames("en", { select: "official" }) as Record<
    string,
    string
  >;

  const items: CountryItem[] = [];

  for (const [cca2Raw, name] of Object.entries(names)) {
    const cca2 = cca2Raw as LibCountryCode;
    try {
      const callingCode = getCountryCallingCode(cca2);
      const label = `+${callingCode}`;
      items.push({ cca2, name, callingCode, label });
    } catch {
      // Some entries may not be supported by libphonenumber-js; skip them.
    }
  }

  items.sort((a, b) => a.name.localeCompare(b.name));
  return items;
}

export default function CountryCodeBox() {
  const { control } = useFormContext<FormValue>();
  const [modalVisible, setModalVisible] = useState(false);
  const [searchText, setSearchText] = useState("");
  const [allCountries] = useState<CountryItem[]>(() => buildCountryItems());
  const [filteredCountries, setFilteredCountries] = useState<CountryItem[]>(
    () => allCountries
  );

  useEffect(() => {
    const q = searchText.trim().toLowerCase();
    const digits = q.replace(/\D/g, "");

    if (!q) {
      setFilteredCountries(allCountries);
      return;
    }

    setFilteredCountries(
      allCountries.filter(
        (c) =>
          c.name.toLowerCase().includes(q) ||
          c.cca2.toLowerCase().includes(q) ||
          (digits ? c.callingCode.includes(digits) : false) ||
          c.label.toLowerCase().includes(q)
      )
    );
  }, [searchText, allCountries]);

  const openModal = () => {
    setModalVisible(true);
  };

  const closeModal = () => {
    setModalVisible(false);
    setSearchText("");
  };

  return (
    <Controller
      name="countryCode"
      control={control}
      defaultValue={"KR" as LibCountryCode}
      render={({ field: { onChange, value } }) => {
        const selected = allCountries.find((c) => c.cca2 === value);
        const display =
          selected?.label ??
          `+${(() => {
            try {
              return getCountryCallingCode((value || "KR") as LibCountryCode);
            } catch {
              return "";
            }
          })()}`;

        return (
          <>
            <Pressable
              accessibilityRole="button"
              onPress={openModal}
              style={styles.box}
            >
              <Text style={styles.boxText} numberOfLines={1}>
                {display}
              </Text>
            </Pressable>

            <Modal
              visible={modalVisible}
              transparent
              animationType="none"
              onRequestClose={closeModal}
            >
              <Pressable style={styles.backdrop} onPress={closeModal} />

              <KeyboardAvoidingView
                behavior={Platform.OS === "ios" ? "padding" : "height"}
                keyboardVerticalOffset={Platform.OS === "ios" ? 20 : 0}
                style={styles.keyboardAvoider}
              >
                <View style={styles.picker}>
                  <View style={styles.handle} />

                  <TextInput
                    value={searchText}
                    onChangeText={setSearchText}
                    placeholder="Search country"
                    placeholderTextColor={colors.GRAY_500}
                    style={styles.search}
                    autoCorrect={false}
                    autoCapitalize="none"
                  />

                  <ScrollView keyboardShouldPersistTaps="handled">
                    {filteredCountries.map((c) => (
                      <Pressable
                        key={c.cca2}
                        style={styles.row}
                        onPress={() => {
                          onChange(c.cca2);
                          closeModal();
                        }}
                      >
                        <Text style={styles.countryName} numberOfLines={1}>
                          {c.name}
                        </Text>
                        <Text style={styles.countryCode} numberOfLines={1}>
                          {c.label}
                        </Text>
                      </Pressable>
                    ))}
                  </ScrollView>
                </View>
              </KeyboardAvoidingView>
            </Modal>
          </>
        );
      }}
    />
  );
}

const styles = StyleSheet.create({
  box: {
    borderWidth: 1,
    borderColor: colors.GRAY_200,
    borderRadius: 10,
    paddingHorizontal: 12,
    paddingVertical: 12,
    minWidth: 92,
    alignItems: "center",
    justifyContent: "center",
    backgroundColor: colors.WHITE,
  },
  boxText: {
    color: colors.BLACK,
    fontSize: 16,
  },

  backdrop: {
    ...StyleSheet.absoluteFillObject,
    backgroundColor: "rgba(0,0,0,0.35)",
  },

  picker: {
    position: "absolute",
    left: 0,
    right: 0,
    bottom: 0,
    height: "50%",
    backgroundColor: colors.WHITE,
    padding: 20,
    borderTopLeftRadius: 16,
    borderTopRightRadius: 16,
  },
  handle: {
    alignSelf: "center",
    width: 44,
    height: 5,
    borderRadius: 999,
    backgroundColor: colors.GRAY_200,
    marginBottom: 12,
  },
  search: {
    borderWidth: 1,
    borderColor: colors.GRAY_200,
    borderRadius: 10,
    paddingHorizontal: 12,
    paddingVertical: 10,
    color: colors.BLACK,
    marginBottom: 12,
  },
  row: {
    paddingVertical: 14,
    borderBottomWidth: StyleSheet.hairlineWidth,
    borderBottomColor: colors.GRAY_200,
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "space-between",
    gap: 12,
  },
  countryName: {
    flex: 1,
    fontSize: 16,
    color: colors.BLACK,
  },
  countryCode: {
    fontSize: 14,
    color: colors.GRAY_700,
  },
  keyboardAvoider: {
    flex: 1,
    justifyContent: "flex-end",
  },
});
