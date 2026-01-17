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
          headerStyle: {
            backgroundColor: colors.GRAY_700,
          },
        }}
      />
      <Stack.Screen
        name="email"
        options={{
          title: "Verify your email",
          headerShown: true,
          headerStyle: {
            backgroundColor: colors.GRAY_700,
          },
        }}
      />
    </Stack>
  );
}
