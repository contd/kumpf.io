<template>
  <v-app id="inspire">
    <v-navigation-drawer
      v-model="drawer"
      app
      :clipped="clipped"
      :floating="floating"
    >
      <v-list dense>
        <v-list-item
          v-for="(item, idx) in items"
          :key="idx"
          :to="item.to"
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title> {{ item.title }} </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar
      app
      clipped-left
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title>{{ title }}</v-toolbar-title>
      <div class="flex-grow-1" />

      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-magnify</v-icon>
      </v-btn>
    </v-app-bar>

    <v-content>
      <v-container
        class="fill-height"
        fluid
      >
        <v-row
          align="center"
          justify="center"
        >
          <v-col class="shrink">
            <v-content>
              <router-view />
            </v-content>
          </v-col>
        </v-row>
      </v-container>
    </v-content>

    <v-footer app>
      <span>&copy; 2019 {{ author }}</span>
    </v-footer>
  </v-app>
</template>

<script>
export default {
  data () {
    return {
      title: 'kumpf.io',
      author: 'Jason Kumpf',
      clipped: true,
      floating: true,
      drawer: false,
      fixed: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'Home',
          to: '/'
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'Inspire',
          to: '/inspire'
        }
      ],
      miniVariant: false,
      right: true,
      rightDrawer: false
    }
  }
}
</script>
