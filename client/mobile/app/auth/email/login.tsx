import { StyleSheet, View } from "react-native";
import { FormProvider, useForm } from "react-hook-form";
import FixedBottomCTA from "@/components/FixedBottomCTA";
import EmailInput from "@/components/EmailInput";
import PasswordInput from "@/components/PasswordInput";
import { useAuth } from "@/hooks/useAuth";
import { colors } from "@/constants";
import { router } from "expo-router";

interface FormValue {
  email: string;
  password: string;
  passwordConfirm: string;
}

export default function LoginScreen() {
  const { emailLoginMutation } = useAuth();

  const emailLoginForm = useForm<FormValue>({
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = (formValues: FormValue) => {
    const { email, password } = formValues;

    emailLoginMutation.mutate(
      {
        email,
        password,
      },
      {
        onSuccess: (data) => {
          if (!data.emailVerified) router.replace("/auth/otp/email");
          if (!data.phoneNumberVerified) router.replace("/auth/phonenumber");
          router.replace("/");
        },
      },
    );
  };
  return (
    <FormProvider {...emailLoginForm}>
      <View style={styles.container}>
        <EmailInput />
        <PasswordInput submitBehavior="submit" />
      </View>
      <FixedBottomCTA
        label="login"
        onPress={emailLoginForm.handleSubmit(onSubmit)}
      />
    </FormProvider>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    margin: 16,
    gap: 16,
    backgroundColor: colors.GRAY_700,
  },
});
