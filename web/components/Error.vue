<script setup lang="ts">
const props = defineProps<{
  message: string
}>()
</script>

<template>
  <div class="error">
    <img src="/pepeSadJam.webp" alt="pepeSadJam">
    <p>{{props.message}}</p>
  </div>
</template>

<style scoped lang="scss">
  .error {
    width: 100%;
    height: 100vh;

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 12px;

    img {
      width: 48px;
      height: 48px;
    }

    p {
      color: var(--foreground-400);

      font: {
        size: 16px;
        weight: 500;
      };

      line-height: 24px;
    }
  }
</style>