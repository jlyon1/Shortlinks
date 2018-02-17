<template>
  <Main></Main>
</template>

<script>

import 'bulma/css/bulma.css'
import Main from '@/components/Main'

export default {
  name: 'App',
  components: {
    Main
  }
}
</script>

<style lang="sass-loader">
@import url('https://fonts.googleapis.com/css?family=Source+Sans+Pro');
$section-padding: 1rem;
$family-source-sans: "Source Sans Pro", sans-serif;
$family-primary: $family-source-sans;
@import '~bulma/bulma';
</style>
