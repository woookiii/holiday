import { useMutation, useQuery } from "@tanstack/react-query";
import Toast from "react-native-toast-message";
import { saveSecureStore } from "@/util/secureStore";
import {
  getMe,
  postEmailLogin,
  postEmailSignup,
  requestEmailOTP,
  requestSMSOTP,
  verifyEmailOTP,
  verifySMSOTP,
} from "@/api/auth";
import { queryKey } from "@/constants";
import { AxiosError } from "axios";

type ResponseError = AxiosError<{
  statusCode: number;
  message: string;
  error: string;
}>;

function useGetMe() {
  const { data } = useQuery({
    queryFn: getMe,
    queryKey: [queryKey.AUTH, queryKey.GET_ME],
  });

  return { data };
}

function useEmailSignup() {
  return useMutation({
    mutationFn: postEmailSignup,
    onSuccess: async (data) => {
      const verificationId = await requestEmailOTP(data);
      saveSecureStore("verificationId", verificationId);
      console.log("success to save verification Id");
    },
    onError: (error: ResponseError) => {
      Toast.show({
        type: "error",
        text1: error.response?.data.message,
      });
    },
  });
}

function useEmailLogin() {
  return useMutation({
    mutationFn: postEmailLogin,
    onSuccess: async (data) => {
      if(!data.emailVerified) {
        const verificationId = await requestEmailOTP(data.id ?? "") //TODO: ask gpt this is okay
        saveSecureStore("verificationId", verificationId)
      }
      if(!data.phoneNumberVerified) {
        saveSecureStore("sessionId", data.sessionId ?? "")
      }
      saveSecureStore("accessToken", data.accessToken ?? "")

    },
    onError: (error: ResponseError) => {
      Toast.show({
        type: "error",
        text1: error.response?.data.message,
      });
    },
  })
}


function useRequestSMSOTP() {
  return useMutation({
    mutationFn: requestSMSOTP,
    onSuccess: (data) => {
      saveSecureStore("verificationId", data?.verificationId);
      console.log("success to save verificationId");
    },
    onError: (error: ResponseError) => {
      Toast.show({
        type: "error",
        text1: error.response?.data.message,
      });
    },
  });
}

function useVerifyEmailOTP() {
  return useMutation({
    mutationFn: verifyEmailOTP,
    onSuccess: (data) => {
      saveSecureStore("sessionId", data?.sessionId)
      console.log("success to save sessionId")
    },
    onError: (error: ResponseError) => {
      Toast.show({
        type: "error",
        text1: error.response?.data.message,
      });
    },
  });
}

function useVerifySMSOTP() {
  return useMutation({
    mutationFn: verifySMSOTP,
    onError: (error: ResponseError) => {
      Toast.show({
        type: "error",
        text1: error.response?.data.message,
      });
    },
  });
}

export function useAuth() {
  // const { data } = useGetMe();
  const emailSignupMutation = useEmailSignup();
  const emailLoginMutation = useEmailLogin();
  const verifyEmailOTPMutation = useVerifyEmailOTP()
  const requestSMSOTPMutation = useRequestSMSOTP();
  const verifySMSOTPMutation = useVerifySMSOTP();

  return {
    auth: {
      //data?.id ||
      id: "",
    },
    emailSignupMutation,
    emailLoginMutation,
    verifyEmailOTPMutation,
    requestSMSOTPMutation,
    verifySMSOTPMutation,
  };
}
