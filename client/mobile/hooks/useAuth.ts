import { useMutation, useQuery } from "@tanstack/react-query";
import Toast from "react-native-toast-message";
import { getSecureStore, saveSecureStore } from "@/util/secureStore";
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
    onError: (error) => {
      Toast.show({
        type: "error",
        text1: error.message,
      });
    },
  });
}

function useEmailLogin() {
  return useMutation({
    mutationFn: postEmailLogin,
    onSuccess: async (data) => {
      if (!data.emailVerified) {
        const { verificationId } = await requestEmailOTP(data.id ?? "");
        console.log(verificationId);
        saveSecureStore("verificationId", verificationId);
        const v = await getSecureStore("verificationId");
        console.log(v);
      }
      if (!data.phoneNumberVerified) {
        saveSecureStore("sessionId", data.sessionId ?? "");
      }
      saveSecureStore("accessToken", data.accessToken ?? "");
    },
    onError: (error) => {
      Toast.show({
        type: "error",
        text1: error.message,
      });
    },
  });
}

function useRequestSMSOTP() {
  return useMutation({
    mutationFn: requestSMSOTP,
    onSuccess: (data) => {
      saveSecureStore("verificationId", data?.verificationId);
      console.log("success to save verificationId");
    },
    onError: (error) => {
      Toast.show({
        type: "error",
        text1: error.message,
      });
    },
  });
}

function useVerifyEmailOTP() {
  return useMutation({
    mutationFn: verifyEmailOTP,
    onSuccess: async (data) => {
      console.log(data?.sessionId);
      saveSecureStore("sessionId", data?.sessionId);
      console.log(await getSecureStore("sessionId"));
    },
    onError: (error) => {
      Toast.show({
        type: "error",
        text1: error.message,
      });
    },
  });
}

function useVerifySMSOTP() {
  return useMutation({
    mutationFn: verifySMSOTP,
    onError: (error) => {
      Toast.show({
        type: "error",
        text1: error.message,
      });
    },
  });
}

export function useAuth() {
  // const { data } = useGetMe();
  const emailSignupMutation = useEmailSignup();
  const emailLoginMutation = useEmailLogin();
  const verifyEmailOTPMutation = useVerifyEmailOTP();
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
