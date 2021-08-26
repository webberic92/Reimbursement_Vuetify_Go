<template>
  <v-layout column align-center class="white--text">
    <v-flex>
      <canvas id="qr"></canvas>
    </v-flex>
    QR code for this App
  </v-layout>
</template>

<script>
import QRCode from 'qrcode';
export default {
  mounted: function(){
    var currentUrl = window.location.origin;
    console.log(currentUrl);
    QRCode.toCanvas(document.getElementById('qr'),
      currentUrl, { toSJISFunc: QRCode.toSJIS }, function (error) {
      if (error) console.error(error)
      console.log('success!')
    })
  }
}
</script>
