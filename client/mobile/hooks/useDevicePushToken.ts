import { useEffect, useState } from "react";
import { Platform } from "react-native";
import * as Device from "expo-device";
import * as Notifications from "expo-notifications";

/**
 * Gets the native device push token.
 *
 * - Android: returns an FCM registration token string.
 * - iOS: returns an APNs device token string.
 *
 * This is the token you use with FCM/APNs directly ("sending notifications custom" guide),
 * NOT the Expo Push Token.
 */
export function useDevicePushToken() {
  const [token, setToken] = useState<string | null>(null);
  const [error, setError] = useState<unknown>(null);

  useEffect(() => {
    let isMounted = true;

    async function run() {
      try {
        if (!Device.isDevice) {
          throw new Error("Must use a physical device to get a push token.");
        }

        // Ask notification permissions (required on iOS; Android 13+ too).
        const { status: existingStatus } =
          await Notifications.getPermissionsAsync();

        let finalStatus = existingStatus;
        if (existingStatus !== "granted") {
          const { status } = await Notifications.requestPermissionsAsync();
          finalStatus = status;
        }

        if (finalStatus !== "granted") {
          throw new Error("Notification permission not granted.");
        }

        // Android: you should set a default notification channel.
        if (Platform.OS === "android") {
          await Notifications.setNotificationChannelAsync("default", {
            name: "default",
            importance: Notifications.AndroidImportance.MAX,
          });
        }

        const t = (await Notifications.getDevicePushTokenAsync()).data;
        if (isMounted) setToken(t);
      } catch (e) {
        if (isMounted) setError(e);
      }
    }

    run();

    return () => {
      isMounted = false;
    };
  }, []);

  return { token, error };
}

