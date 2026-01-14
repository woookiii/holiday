import { Stack } from "expo-router";

export default function EmailLayout() {
  return (
    <Stack>
      <Stack.Screen
        name="signup"
        options={{
          title: "Sign up with your email",
          headerShown: true,
        }}
      />
      <Stack.Screen
        name="login"
        options={{
          title: "Login with your email",
          headerShown: true,
        }}
      />
    </Stack>
  );
}
