import { Stack } from "expo-router";

export default function AuthLayout() {

  return <Stack>
    <Stack.Screen
      name="index"
      options={{
        title: "",
        headerShown: false
      }}
    />
    <Stack.Screen
      name="phonenumber"
      options={{
        title: "Verify your phone number",
        headerShown: true
      }}
    />
    <Stack.Screen
      name="otp/sms"
      options={{
        title: "Verify your OTP",
        headerShown: true
      }}
    />
  </Stack>;
}