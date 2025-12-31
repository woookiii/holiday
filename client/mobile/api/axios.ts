import axios from "axios";
import { getSecureStore, saveSecureStore } from "@/util/secureStore";
import { Platform } from "react-native";

const baseUrl = {
  android: "http://10.0.2.2:8080",
  ios: "http://localhost:8080"
};

const axiosInstance = axios.create({
  baseURL: Platform.OS === "ios" ? baseUrl.ios : baseUrl.android,
  withCredentials: true,
  headers: {
    "Content-Type": "application/json"
  }
});

//circular import warning is annoying and it is here
async function refreshAccessToken() {
  try {
    const res = await axiosInstance.post("/auth/refresh-token");
    return res.data;
  } catch (err: any) {
    const message = err.response?.data?.message || "Failed to refresh access token";
    throw new Error(message);
  }
}

axiosInstance.interceptors.request.use((config) => {
  const token = getSecureStore("accessToken");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
});

axiosInstance.interceptors.response.use((res) => res,
  async (error) => {
    if (error.response?.status) {
      const originalRequest = error.config;

      if (error.response?.status === 401 && !originalRequest._retry && !originalRequest.url.includes("/refresh-token")) {
        originalRequest._retry = true;
        try {
          const { newToken } = await refreshAccessToken();
          await saveSecureStore("accessToken", newToken);
          originalRequest.headers.Authorization = `Bearer ${newToken}`;
          return axiosInstance(originalRequest);
        } catch (err) {
          console.error("Refresh token failed", err);
        }
      }
      return Promise.reject(error);
    }
  }
);

export { baseUrl, axiosInstance };