import { axiosInstance } from "@/api/axios";
import { Token } from "@/types";


async function requestSmsOtpToFirebase(phoneNumber: string) {
  return Promise.resolve("good");
}

async function postSmsOtpToFirebase(otp: string) {
}

async function postFirebaseTokenToServer(firebaseToken: string) {
  const { data } = await axiosInstance.post("/auth/firebase-token", firebaseToken);
  return data;
}

async function getMe() {
  const { data } = await axiosInstance.get("/auth/me");

  return data;
}

export {
  requestSmsOtpToFirebase,
  postSmsOtpToFirebase,
  postFirebaseTokenToServer,
  getMe
};
