import { useMutation, useQuery } from "@tanstack/react-query";
import Toast from "react-native-toast-message";
import { saveSecureStore } from "@/util/secureStore";
import { FirebaseError } from "@firebase/util";
import {
  getMe,
  postFirebaseTokenToServer,
  postSmsOtpToFirebase,
  requestSmsOtpToFirebase
} from "@/api/auth";
import { queryKey } from "@/constants";

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
      await saveSecureStore("verificationId", confirmationResult.verificationId);
      console.log("success to save verification Id")
    },
    onError: (error: FirebaseError) => {
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
    onError: (error: FirebaseError) => {
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
