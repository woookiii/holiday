import { Stack } from 'expo-router';
import 'react-native-reanimated';
import { FIREBASE_AUTH } from "@/firebaseConfig";
import { QueryClientProvider } from "@tanstack/react-query";
import queryClient from "@/api/queryClient";


export default function RootLayout() {

  FIREBASE_AUTH.useDeviceLanguage()
  return (
    <QueryClientProvider client={queryClient}>
      <RootNavigator/>
    </QueryClientProvider>
  );
}

function RootNavigator() {
  return <Stack>
    <Stack.Screen name="(init)" options={{headerShown: false}}/>
    <Stack.Screen name="(tabs)" options={{ headerShown: false }}/>
    <Stack.Screen name="auth" options={{headerShown: false}}/>
  </Stack>
}
