<script lang="ts" setup>
</script>

<template>
  <a href="https://discord.com/oauth2/authorize?client_id=1207683042680897629&response_type=code&redirect_uri=http%3A%2F%2Flocalhost%3A3000%2Ftoken&scope=identify+guilds+guilds.members.read">auth</a>
</template>

<style lang="scss" scoped>
  .wrapper {
    padding: 24px;

    h3 {
      color: var(--foreground-500);

      font-size: 24px;

      line-height: 48px;
    }

    pre {
      color: var(--foreground-400);
    }
  }
</style>