import { useMutation, useQuery } from "@tanstack/react-query";
import Toast from "react-native-toast-message";
import { saveSecureStore } from "@/util/secureStore";
import {
  getMe,
  postFirebaseTokenToServer,
  postSmsOtpToFirebase,
  requestSmsOtpToFirebase
} from "@/api/auth";
import { queryKey } from "@/constants";
import { FirebaseAuthTypes } from "@react-native-firebase/auth";

function useGetMe() {
  const { data } = useQuery({
    queryFn: getMe,
    queryKey: [queryKey.AUTH, queryKey.GET_ME]
  });

  return { data };
}

function useRequestSmsOtp() {
  return useMutation({
    mutationFn: requestSmsOtpToFirebase,
    onSuccess: async (confirmationResult) => {
      const verificationId = confirmationResult.verificationId;

      if (!verificationId) {
        Toast.show({
          type: "error",
          text1: "fail to request sms"
        });
        return;
      }
      await saveSecureStore("verificationId", verificationId);
      console.log("success to save verification Id");
    },
    onError: (error: FirebaseAuthTypes.NativeFirebaseAuthError) => {
      Toast.show({
        type: "error",
        text1: error.message
      });
    }
  });
}

function usePostSmsOtp() {
  return useMutation({
    mutationFn: postSmsOtpToFirebase,
    onSuccess: async (data) => {
      const firebaseToken = await data.user.getIdToken();
      const { accessToken } = await postFirebaseTokenToServer(firebaseToken);
      await saveSecureStore("accessToken", accessToken);
    },
    onError: (error: FirebaseAuthTypes.NativeFirebaseAuthError) => {
      Toast.show({
        type: "error",
        text1: error.message
      });
    }
  });
}

function useAuth() {
  // const { data } = useGetMe();
  const requestSmsOtpMutation = useRequestSmsOtp();
  const postSmsOtpMutation = usePostSmsOtp();


  return {
    auth: {
      id: //data?.id ||
        ""
    },
    requestSmsOtpMutation,
    postSmsOtpMutation
  };
}

export { useAuth };
