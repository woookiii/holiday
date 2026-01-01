import { useMutation, useQuery } from "@tanstack/react-query";
import Toast from "react-native-toast-message";
import { saveSecureStore } from "@/util/secureStore";
import { getMe, requestSmsOtp, verifySmsOtp } from "@/api/auth";
import { queryKey } from "@/constants";

function useGetMe() {
  const { data } = useQuery({
    queryFn: getMe,
    queryKey: [queryKey.AUTH, queryKey.GET_ME],
  });

  return { data };
}

function useRequestSmsOtp() {
  return useMutation({
    mutationFn: requestSmsOtp,
    onSuccess: (data) => {
      saveSecureStore("verificationId", data?.verificationId)
      console.log("success to save verification Id");
    },
  });
}

function usePostSmsOtp() {
  return useMutation({
    mutationFn: verifySmsOtp,
    onSuccess: (data) => {},
    onError: (error) => {
      Toast.show({
        type: "error",
        text1: error.message,
      });
    },
  });
}

function useAuth() {
  // const { data } = useGetMe();
  const requestSmsOtpMutation = useRequestSmsOtp();
  const postSmsOtpMutation = usePostSmsOtp();

  return {
    auth: {
      //data?.id ||
      id: "",
    },
    requestSmsOtpMutation,
    postSmsOtpMutation,
  };
}

export { useAuth };
