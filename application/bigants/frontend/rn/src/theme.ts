import { TextStyle, ViewStyle, Platform } from 'react-native';

const colors = {
  primary: {
    black: '#121314',
    blue: '#0084F4',
    red: '#FC0000',
    white: '#ffffff',
  },
  text: {
    content: '#4A4A54',
    light: '#757C85',
  },
  background: {
    box: '#EFEFEF',
    roundButton: '#ffffff'
  },
  border: {
    box: '#EFEFEF'
  }
};

const fontSize = {
  xxl: 32,
  xl: 24,
  l: 18,
  m: 16,
  s: 14,
  xs: 12,
}

const MU = {
  V1: 8,
  V2: 16,
  V3: 24,
  V4: 32,
  V5: 40,
  V6: 48,

  H1: 10,
  H2: 20,
  H3: 30,
}

const xxlbText: TextStyle = {
  fontSize: fontSize.xxl,
  fontWeight: 'bold',
  color: colors.primary.black,
  lineHeight: 40,
  textAlignVertical: 'center',
};

const xlwText: TextStyle = {
  fontSize: fontSize.xl,
  fontWeight: 'normal',
  color: colors.primary.white,
  lineHeight: 40,
  textAlignVertical: 'center',
};

const xlwbText: TextStyle = {
  fontSize: fontSize.xl,
  fontWeight: 'bold',
  color: colors.primary.white,
  lineHeight: 40,
  textAlignVertical: 'center',
};

const xlbbText: TextStyle = {
  fontSize: fontSize.xl,
  fontWeight: 'bold',
  color: colors.primary.black,
  lineHeight: 40,
  textAlignVertical: 'center',
};

const lnText: TextStyle = {
  fontSize: fontSize.l,
  fontWeight: 'normal',
  color: colors.primary.black,
  lineHeight: 28,
  letterSpacing: -0.5,
  textAlignVertical: 'center',
};

const lbbText: TextStyle = {
  fontSize: fontSize.l,
  fontWeight: 'bold',
  color: colors.primary.black,
  lineHeight: 28,
  letterSpacing: -0.5,
  textAlignVertical: 'center',
};

const lbblueText: TextStyle = {
  fontSize: fontSize.l,
  fontWeight: 'bold',
  color: colors.primary.blue,
  lineHeight: 28,
  letterSpacing: -0.5,
  textAlignVertical: 'center',
};

const mnText: TextStyle = {
  fontSize: fontSize.m,
  fontWeight: 'normal',
  color: colors.text.content,
  lineHeight: 22,
  textAlignVertical: 'center',
};

const mbbText: TextStyle = {
  fontSize: fontSize.m,
  fontWeight: 'bold',
  color: colors.primary.black,
  lineHeight: 22,
  textAlignVertical: 'center',
};

const mbblueText: TextStyle = {
  fontSize: fontSize.m,
  fontWeight: 'bold',
  color: colors.primary.blue,
  lineHeight: 22,
  textAlignVertical: 'center',
};

const sbText: TextStyle = {
  fontSize: fontSize.s,
  fontWeight: 'bold',
  color: colors.text.content,
  lineHeight: 20,
  textAlignVertical: 'center',
}

const sbbText: TextStyle = {
  fontSize: fontSize.s,
  fontWeight: 'bold',
  color: colors.primary.black,
  lineHeight: 20,
  textAlignVertical: 'center',
}

const snText: TextStyle = {
  fontSize: fontSize.s,
  fontWeight: 'normal',
  color: colors.text.content,
  lineHeight: 20,
  textAlignVertical: 'center',
}

const xsnText: TextStyle = {
  fontSize: fontSize.xs,
  fontWeight: 'normal',
  color: colors.text.content,
  lineHeight: 20,
  textAlignVertical: 'center',
}

const boxShadow: ViewStyle = {
  ...Platform.select({
    ios: {
      shadowColor: colors.primary.black,
      shadowOffset: { width: 0, height: 1 },
      shadowOpacity: 0.3,
      shadowRadius: 0.8 * 4,
    },
    android: {
      elevation: 4,
    }
  })
};

export const theme = {
  colors,
  fontSize,
  MU,
  boxShadow,
  xxlbText,
  xlwbText,
  xlwText,
  xlbbText,
  lbbText,
  lbblueText,
  lnText,
  mbbText,
  mbblueText,
  mnText,
  sbText,
  sbbText,
  snText,
  xsnText,
}