import { Stack, useFocusEffect } from "expo-router";
import { Pressable } from "react-native";

export default function InitLayout() {

  return <Stack>
    <Stack.Screen
      name="index"
      options={{
        headerShown: false
      }}
    />
  </Stack>;
}