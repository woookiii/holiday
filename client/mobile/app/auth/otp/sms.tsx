import { StyleSheet, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { CountryCode } from "libphonenumber-js";
import InputField from "@/components/InputField";
import { colors } from "@/constants";
import FixedBottomCTA from "@/components/FixedBottomCTA";
import { FormProvider, useForm } from "react-hook-form";
import SmsCodeInput from "@/components/SmsCodeInput";

interface FormValue {
  smsCode: string;
}

export default function SmsScreen() {
  const profileForm = useForm<FormValue>({
    defaultValues: {
      smsCode: "",
    },
  });

  return (
    <FormProvider {...profileForm}>
      <SafeAreaView style={styles.container}>
        <View style={styles.content}>
          <SmsCodeInput />
        </View>
        <FixedBottomCTA label="Confirm" />
      </SafeAreaView>
    </FormProvider>
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
    paddingHorizontal: 100,
    paddingTop: 120,
  },
});
