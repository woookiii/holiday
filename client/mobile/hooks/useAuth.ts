import { useMutation, useQuery } from "@tanstack/react-query";
import { router } from "expo-router";
import Toast from "react-native-toast-message";
import { saveSecureStore } from "@/util/secureStore";
import { FirebaseError } from "@firebase/util";
import { getMe, postFirebaseTokenToServer, postPhoneOtpToFirebase, requestOtpToFirebase } from "@/api/auth";
import { queryKey } from "@/constants";

function useGetMe() {
  const {data} = useQuery({
    queryFn: getMe,
    queryKey: [queryKey.AUTH, queryKey.GET_ME]
  });

  return { data }
}

function useRequestOtp() {
  return useMutation({
    mutationFn: requestOtpToFirebase,
    onSuccess: async (confirmationResult) => {
      await saveSecureStore("verificationId", confirmationResult.verificationId);
      router.push("/auth/otp");
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
      router.replace("/")
    }
  });
}

function useAuth() {
  const requestOtpMutation = useRequestOtp();
  const postPhoneOtpMutation = usePostPhoneOtp();


  return {
    auth: {
      id: null
    },
    requestOtpMutation,
    postPhoneOtpMutation,
  };
}

export { useAuth };
