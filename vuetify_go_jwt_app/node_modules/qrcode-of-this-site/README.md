# QRcode of this site

<img src="/img/IMG_0228.PNG" width="50%"/>

## Install 
```
yarn add qrcode-of-this-site
#or
npm install qrcode-of-this-site
```

## How to use
Import **QRcode**, then put **<QRcode/>** on your Vue template as follows:

```vue:
<template>
  <QRcode/>
<template>

<script>
import {QRcode} from 'qrcode-of-this-site'
export default {
  components: {QRcode},
}
</script>
```

A sample of usage in the demo site code is [here](https://github.com/UedaTakeyuki/vue-faui-user-fe-sample/blob/275e6752883ea11400814995c5b0830a3227f84e/src/components/Navbar.vue#L36).

## Demo
Demo site is available at [here](https://vue-faui-user-fe-sample.uedasoft.com/).
You can see QR code of that site in the top of the navigation drawer.

## Q&A
Github [issues](https://github.com/UedaTakeyuki/qrcode-of-this-site/issues) are available. Any questions, suggestions, reports are welcome!

## Author
[Dr. Takeyuki UEDA](https://atelierueda.uedasoft.com/)

## History
- 1.0.4 2021.04.01 first version.
- 1.1.1 2021.06.23 decrease package size