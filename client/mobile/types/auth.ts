export interface Token {
  accessToken: string;
  refreshToken: string;
}

export interface emailSignReq {
  email: string;
  password: string;
}

export interface verifyEmailOTPReq {
  otp: string
  verificationId: string
}

export interface verifySMSOTPReq {
  otp: string;
  verificationId: string;
  sessionId: string | null;
}


export interface EmailLogiResp {
  emailVerified: boolean;
  phoneNumberVerified: boolean;
  id: string | null;
  sessionId: string | null;
  accessToken: string | null;
}


