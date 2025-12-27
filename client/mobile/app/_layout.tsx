import { Stack } from 'expo-router';
import 'react-native-reanimated';
import { FIREBASE_AUTH } from "@/firebaseConfig";


export default function RootLayout() {

  FIREBASE_AUTH.useDeviceLanguage()
  return (
      <Stack>
        <Stack.Screen name="(init)" options={{headerShown: false}}/>
        <Stack.Screen name="(tabs)" options={{ headerShown: false }}/>
        <Stack.Screen name="auth" options={{headerShown: false}}/>
      </Stack>
  );
}
