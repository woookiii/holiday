import { Stack } from "expo-router";

export default function AuthLayout() {
  return <Stack>
    <Stack.Screen
      name="index"
      options={{
        title:"로그인",
        headerShown: false
      }}
    />
  </Stack>;
}