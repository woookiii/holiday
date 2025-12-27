import { useMutation } from "@tanstack/react-query";
import { router } from "expo-router";
import Toast from "react-native-toast-message";
import { saveSecureStore } from "@/util/secureStore";
import { FirebaseError } from "@firebase/util";
import { postFirebaseTokenToServer, postPhoneOtpToFirebase, requestOtpToFirebase } from "@/api/auth";


function useRequestOtp() {
  return useMutation({
    mutationFn: requestOtpToFirebase,
    onSuccess: async (confirmationResult) => {
      await saveSecureStore("verificationId", confirmationResult.verificationId);
      router.replace("/auth/otp");
    },
    onError: (error: FirebaseError) => {
      Toast.show({
        type: "error",
        text1: error.message
      });
    }
  });
}

function usePostPhoneOtp() {
  return useMutation({
    mutationFn: postPhoneOtpToFirebase,
    onSuccess: async (data) => {
      const firebaseToken = await data.user.getIdToken()
      const { accessToken } = await postFirebaseTokenToServer(firebaseToken);
      await saveSecureStore("accessToken", accessToken);
    }
  });
}

function useAuth() {


  return {
    auth: {
      id: null
    }
  };
}

export { useAuth };
