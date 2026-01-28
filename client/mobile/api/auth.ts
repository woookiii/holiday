import { axiosInstance } from "@/api/axios";
import {
  EmailLoginResp,
  emailSignReq,
  signInWithAppleReq,
  signInWithAppleResp,
  verifyEmailOTPReq,
  verifyEmailOTPResp,
  verifySMSOTPReq,
  verifySMSOTPResp,
} from "@/types/auth";

export async function postEmailSignup(body: emailSignReq): Promise<string> {
  console.log("post email sign up");
  const { data } = await axiosInstance.post("/auth/email/create", body);
  return data;
}

export async function postEmailLogin(
  body: emailSignReq,
): Promise<EmailLoginResp> {
  const { data } = await axiosInstance.post("/auth/email/login", body);
  console.log(data);
  return data;
}

export async function requestEmailOTP(
  id: string,
): Promise<{ verificationId: string }> {
  const { data } = await axiosInstance.post("/auth/email/otp/send", { id });
  return data;
}

export async function requestSMSOTP(phoneNumber: string) {
  const { data } = await axiosInstance.post("/auth/sms/otp/send", {
    phoneNumber,
  });
  return data;
}

export async function verifyEmailOTP(
  body: verifyEmailOTPReq,
): Promise<verifyEmailOTPResp> {
  const { data } = await axiosInstance.post("/auth/email/otp/verify", body);
  return data;
}

export async function verifySMSOTP(
  body: verifySMSOTPReq,
): Promise<verifySMSOTPResp> {
  const { data } = await axiosInstance.post("/auth/sms/otp/verify", body);
  return data;
}

export async function signInWithApple(
  body: signInWithAppleReq,
): Promise<signInWithAppleResp> {
  const { data } = await axiosInstance.post("/auth/apple", body);
  return data;
}

export async function getMe() {
  const { data } = await axiosInstance.get("/auth/me");
  return data;
}
