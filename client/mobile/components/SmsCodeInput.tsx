import { StyleSheet, View } from "react-native";
import { Controller, useFormContext } from "react-hook-form";
import InputField from "@/components/InputField";

export default function SmsCodeInput() {
  const { control } = useFormContext();

  return (
    <Controller
      name="smsCode"
      control={control}
      rules={{
        validate: (data: string) => {
          if (data.length <= 0) {
            return "Put your Code before Verify";
          }
        },
      }}
      render={({ field: { onChange, value }, fieldState: { error } }) => (
        <InputField
          autoFocus
          label="sms code"
          placeholder="0000000"
          returnKeyType="done"
          submitBehavior="submit"
          value={value}
          onChangeText={onChange}
          error={error?.message}
        />
      )}
    />
  );
}
