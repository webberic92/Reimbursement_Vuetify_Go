<template>
  <nav>
    <v-app-bar>
      <v-app-bar-nav-icon class="grey--text" @click="drawer = !drawer"></v-app-bar-nav-icon>

      <slot name="title">
        <v-toolbar-title class="grey--text">
          <span>{{titleStr}}</span>
        </v-toolbar-title>
      </slot>
      <slot name="menu">
        <v-spacer></v-spacer>

        <!-- dropdown menu -->
        <v-menu offset-y>
          <template v-slot:activator="{ on }">
          <!-- <v-btn text slot="activator"> -->
          <v-btn text v-on="on">
            <v-icon left>expand_more</v-icon>
            <span>Menu</span>
          </v-btn>
          </template>
          <v-list>
            <!-- v-list-tile is changed to v-list-item -->
            <v-list-item v-for="link in links" :key="link.text" router :to="link.route">
              <v-list-item-title>{{ link.text }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </slot>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app class="primary">
      <v-list>
        <QRcode/>
        <!-- v-list-tile is changed to v-list-item -->
        <v-list-item v-for="link in links" :key="link.text" router :to="link.route">
          <v-list-item-action>
            <v-icon class="white--text">{{ link.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title class="white--text">{{ link.text }}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </nav>
</template>


<script>
import {QRcode} from 'qrcode-of-this-site'
//import {AccountRoutes, PurchaseRoutes} from 'vue-faui-user-fe/'
export default {
  props: {
    titleStr: String,
    links: Array,
  },
  components: {QRcode},
  data: () => ({
    drawer: false,
/*    links: [
      { icon: 'home', text: 'Home', route: '/'},
      AccountRoutes.menuItem,
      PurchaseRoutes.menuItem,
    ]*/
  }),
}
</script>