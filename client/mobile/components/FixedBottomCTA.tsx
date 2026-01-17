import { colors } from "@/constants";
import React from "react";
import { PressableProps, StyleSheet, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import CustomButton from "./CustomButton";

interface FixedBottomCTAProps extends Omit<PressableProps, "style"> {
  label: string;
}

function FixedBottomCTA({ label, ...props }: FixedBottomCTAProps) {
  const inset = useSafeAreaInsets();

  return (
    <View style={[styles.fixed, { paddingBottom: inset.bottom || 12 }]}>
      <CustomButton variant="outlined" label={label} {...props} />
    </View>
  );
}

const styles = StyleSheet.create({
  fixed: {
    position: "absolute",
    bottom: 10,
    width: "100%",
    borderTopWidth: StyleSheet.hairlineWidth,
    borderTopColor: colors.GRAY_700,
    paddingTop: 12,
    paddingHorizontal: 16,
  },
});

export default FixedBottomCTA;
