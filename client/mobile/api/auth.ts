import { axiosInstance } from "@/api/axios";
import { Token } from "@/types";

interface verifySmsOtpReq {
  smsCode: string;
  verificationId: string;
}

async function requestSmsOtp(phoneNumber: string) {
  const { data } = await axiosInstance.post("/auth/sms/otp", phoneNumber);
  return data;
}

async function verifySmsOtp(verifySmsOtpReq: verifySmsOtpReq) {
  const { data } = await axiosInstance.post("/auth/sms/otp", verifySmsOtpReq);
  return data;
}

async function getMe() {
  const { data } = await axiosInstance.get("/auth/me");
  return data;
}

export { requestSmsOtp, verifySmsOtp, getMe };
