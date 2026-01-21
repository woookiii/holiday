import { Stack } from "expo-router";
import { colors } from "@/constants";

export default function OTPLayout() {
  return (
    <Stack>
      <Stack.Screen
        name="sms"
        options={{
          title: "Verify your phone number",
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="email"
        options={{
          title: "Verify your email",
          headerShown: true,
        }}
      />
    </Stack>
  );
}
