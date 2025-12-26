import { Stack } from "expo-router";

export default function InitLayout() {
  return <Stack>
    <Stack.Screen
      name="index"
      options={{
        title:"home",
        headerShown: false
      }}
    />
  </Stack>;
}