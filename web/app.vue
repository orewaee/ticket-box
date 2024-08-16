<script setup lang="ts">

</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<style lang="scss">
@import url("https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&display=swap");

* {
  box-sizing: border-box;

  margin: 0;
  padding: 0;
}

::selection {
  color: var(--background-500);

  background-color: var(--accent);
}

body {
  --background-500: #12131A;
  --background-400: #14151D;
  --background-300: #161823;
  --background-200: #1A1C2A;
  --background-100: #202231;

  --background-500-40: rgba(18, 19, 26, .4);
  --background-300-40: rgba(22, 24, 35, .4);

  --foreground-500: #EAEBED;
  --foreground-400: #D8D9DE;
  --foreground-300: #BABCC7;
  --foreground-200: #999DB5;
  --foreground-100: #73799F;

  --accent: #5865F2;

  --red-500-100: #ED4245;
  --red-500-40: rgba(237, 66, 69, .4);
  --red-500-20: rgba(237, 66, 69, .2);

  background-color: var(--background-500);

  h1, h2, h3, h4, h5, h6, p, span, a, small, strong, input, button {
    font-family: Inter, sans-serif;
  }

  h1 {
    font: {
      size: 32px;
      weight: 500;
    };

    line-height: 48px;
  }
}
</style>
