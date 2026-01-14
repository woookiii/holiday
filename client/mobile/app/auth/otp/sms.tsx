import { StyleSheet, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { colors } from "@/constants";
import FixedBottomCTA from "@/components/FixedBottomCTA";
import { FormProvider, useForm } from "react-hook-form";
import SmsCodeInput from "@/components/SmsCodeInput";
import { useAuth } from "@/hooks/useAuth";
import { router } from "expo-router";
import { getSecureStore } from "@/util/secureStore";
import Toast from "react-native-toast-message";

interface FormValue {
  smsCode: string;
}

export default function OTPSmsScreen() {
  const { postSmsOtpMutation } = useAuth();
  const profileForm = useForm<FormValue>({
    defaultValues: {
      smsCode: "",
    },
  });

  const onSubmit = async (formValue: FormValue) => {
    const { smsCode } = formValue;
    const verificationId = await getSecureStore("verificationId");
    if (!verificationId) {
      console.error("fail to get verification Id");
      Toast.show({
        type: "error",
        text1: "you can't send code",
      });
      return;
    }
    console.log("execute post sms otp mutate");
    postSmsOtpMutation.mutate(
      { smsCode, verificationId },
      {
        onSuccess: () => router.replace("/home"),
      },
    );
  };

  return (
    <FormProvider {...profileForm}>
      <SafeAreaView style={styles.container}>
        <View style={styles.content}>
          <SmsCodeInput />
        </View>
        <FixedBottomCTA
          label="Confirm"
          onPress={profileForm.handleSubmit(onSubmit)}
        />
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
