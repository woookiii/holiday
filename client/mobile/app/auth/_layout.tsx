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
        title: "Verify your phone number",
        headerShown: true
      }}
    />
    <Stack.Screen
      name="otp/email"
      options={{
        title: "Verify your email",
        headerShown: true
      }}
    />
    <Stack.Screen
      name="email/signup"
      options={{
        title: "Sign up with your email",
        headerShown: true
      }}
    />
    <Stack.Screen
      name="email/login"
      options={{
        title: "Login with your email",
        headerShown: true
      }}
    />
  </Stack>;
}