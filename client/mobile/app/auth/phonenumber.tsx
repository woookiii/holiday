import { StyleSheet, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { colors } from "@/constants";
import FixedBottomCTA from "@/components/FixedBottomCTA";
import CountryCodeBox from "@/components/CountryCodeBox";
import PhoneNumberInput from "@/components/PhoneNumberInput";
import { FormProvider, useForm } from "react-hook-form";
import {
  CountryCode,
  isValidPhoneNumber,
  parsePhoneNumberFromString,
} from "libphonenumber-js";
import Toast from "react-native-toast-message";
import { getAuth, signInWithPhoneNumber } from "firebase/auth";
import { FIREBASE_AUTH, RECAPTCHA } from "@/firebaseConfig";

interface FormValue {
  countryCode: CountryCode;
  phoneNumber: string;
}

export default function PhonenumberScreen() {
  const phoneNumberForm = useForm<FormValue>({
    defaultValues: {
      countryCode: "KR",
      phoneNumber: "",
    },
  });

  const onSubmit = (formValues: FormValue) => {
    const { countryCode, phoneNumber } = formValues;

    const digitsOnly = phoneNumber.replace(/[^\d]/g, "");

    if (!isValidPhoneNumber(digitsOnly, countryCode)) {
      Toast.show({
        type: "error",
        text1: "Invalid phone number",
      });
      return;
    }

    const parsed = parsePhoneNumberFromString(digitsOnly, countryCode);
    const wholeNumber = parsed?.number;

    if (!wholeNumber) {
      Toast.show({
        type: "error",
        text1: "Invalid phone number",
      });
      return;
    }

    // wholeNumber is E.164 (e.g. +821012345678). Use this for Firebase.
    // TODO: connect firebase
    signInWithPhoneNumber(FIREBASE_AUTH, wholeNumber, RECAPTCHA)
  };

  return (
    <SafeAreaView style={styles.container}>
      <FormProvider {...phoneNumberForm}>
        <View style={styles.content}>
          <View style={styles.phoneRow}>
            <CountryCodeBox />
            <PhoneNumberInput />
          </View>
        </View>

        <FixedBottomCTA
          label={"Send Code"}
          onPress={phoneNumberForm.handleSubmit(onSubmit)}
        />
      </FormProvider>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: colors.WHITE,
  },
  content: {
    flex: 1,
    width: "100%",
    paddingHorizontal: 20,
    paddingTop: 120,
  },
  phoneRow: {
    flexDirection: "row",
    width: "100%",
    alignItems: "center",
    gap: 10,
  },
});