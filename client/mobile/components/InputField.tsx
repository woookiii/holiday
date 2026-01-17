import { colors } from "@/constants";
import { ForwardedRef, forwardRef, ReactNode } from "react";
import {
  StyleSheet,
  Text,
  TextInput,
  TextInputProps,
  View,
} from "react-native";

interface InputFieldProps extends TextInputProps {
  label?: string;
  variant?: "filled" | "standard" | "outlined";
  error?: string;
  rightChild?: ReactNode;
}

function InputField(
  {
    label,
    variant = "filled",
    error = "",
    rightChild = null,
    ...props
  }: InputFieldProps,
  ref?: ForwardedRef<TextInput>,
) {
  return (
    <View>
      {label && <Text style={s.label}>{label}</Text>}
      <View
        style={[
          s.container,
          s[variant],
          props.multiline && s.multiline,
          Boolean(error) && s.inputError,
        ]}
      >
        <TextInput
          ref={ref}
          style={[s.input, s[`${variant}Text`]]}
          autoCapitalize="none"
          placeholderTextColor={colors.GRAY_400}
          spellCheck={false}
          autoCorrect={false}
          {...props}
        />
        {rightChild}
      </View>
      {Boolean(error) && <Text style={s.error}>{error}</Text>}
    </View>
  );
}

const s = StyleSheet.create({
  label: {
    fontSize: 12,
    color: colors.GRAY_700,
    marginBottom: 5,
  },
  container: {
    height: 44,
    borderRadius: 8,
    paddingHorizontal: 10,
    justifyContent: "center",
    alignItems: "center",
    flexDirection: "row",
    alignSelf: "stretch",
    width: "100%",
  },
  filled: {
    backgroundColor: colors.GRAY_100,
  },
  standard: {
    borderWidth: 1,
    borderColor: colors.GRAY_200,
    backgroundColor: colors.WHITE,
  },
  outlined: {
    borderWidth: 1,
    borderColor: colors.GREEN_600,
  },
  standardText: {
    color: colors.BLACK,
  },
  outlinedText: {
    color: colors.BLACK,
    fontWeight: "bold",
  },
  filledText: {},
  input: {
    fontSize: 16,
    padding: 0,
    flex: 1,
  },
  error: {
    fontSize: 12,
    marginTop: 5,
    color: colors.RED_500,
  },
  inputError: {
    backgroundColor: colors.RED_100,
  },
  multiline: {
    alignItems: "flex-start",
    paddingVertical: 10,
    height: 200,
  },
});

export default forwardRef(InputField);
