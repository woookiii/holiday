import { axiosInstance } from "@/api/axios";

interface verifySmsOtpReq {
  smsCode: string;
  verificationId: string;
}

async function requestSmsOtp(phoneNumber: string) {
  const { data } = await axiosInstance.post("/auth/sms/otp/send", {
    phoneNumber,
  });
  return data;
}

async function verifySmsOtp(verifySmsOtpReq: verifySmsOtpReq) {
  const { data } = await axiosInstance.post(
    "/auth/sms/otp/verify",
    verifySmsOtpReq,
  );
  return data;
}

async function getMe() {
  const { data } = await axiosInstance.get("/auth/me");
  return data;
}

export { requestSmsOtp, verifySmsOtp, getMe };
