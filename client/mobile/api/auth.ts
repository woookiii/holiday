import { axiosInstance } from "@/api/axios";
import {
  EmailLogiResp,
  emailSignReq,
  verifyEmailOTPReq,
  verifySMSOTPReq,
} from "@/types/auth";



export async function postEmailSignup(body: emailSignReq) : Promise<string> {
  const { data } = await axiosInstance.post("/auth/email/signup", body);
  return data
}

export async function postEmailLogin(
  body: emailSignReq,
): Promise<EmailLogiResp> {
  const { data } = await axiosInstance.post("/auth/email/login", body);
  return data;
}

export async function requestEmailOTP(id: string) : Promise<string>{
  const { data } = await axiosInstance.post("/auth/email/otp/send", id)
  return data
}

export async function requestSMSOTP(phoneNumber: string) {
  const { data } = await axiosInstance.post("/auth/sms/otp/send", {
    phoneNumber,
  });
  return data;
}

export async function verifyEmailOTP(body: verifyEmailOTPReq) {
  const { data } = await axiosInstance.post("/auth/email/otp/verify", body);
  return data;
}

export async function verifySMSOTP(verifySMSOTPReq: verifySMSOTPReq) {
  const { data } = await axiosInstance.post(
    "/auth/sms/otp/verify",
    verifySMSOTPReq,
  );
  return data;
}

export async function getMe() {
  const { data } = await axiosInstance.get("/auth/me");
  return data;
}
