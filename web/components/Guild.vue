<script setup lang="ts">
const props = defineProps<{
  id: string
  name: string
  iconUrl: string
  withBot: boolean
}>()
</script>

<template>
  <div class="guild">
    <img draggable="false" :src="props.iconUrl" alt="">
    <div class="info">
      <div class="name">
        <h4>{{props.name}}</h4>
        <svg v-if="props.withBot" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="var(--accent)" class="icon icon-tabler icons-tabler-filled icon-tabler-square-check">
          <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
          <path d="M18.333 2c1.96 0 3.56 1.537 3.662 3.472l.005 .195v12.666c0 1.96 -1.537 3.56 -3.472 3.662l-.195 .005h-12.666a3.667 3.667 0 0 1 -3.662 -3.472l-.005 -.195v-12.666c0 -1.96 1.537 -3.56 3.472 -3.662l.195 -.005h12.666zm-2.626 7.293a1 1 0 0 0 -1.414 0l-3.293 3.292l-1.293 -1.292l-.094 -.083a1 1 0 0 0 -1.32 1.497l2 2l.094 .083a1 1 0 0 0 1.32 -.083l4 -4l.083 -.094a1 1 0 0 0 -.083 -1.32z" />
        </svg>
      </div>
      <h4 class="id">{{props.id}}</h4>
    </div>
  </div>
</template>

<style scoped lang="scss">
  .guild {
    display: flex;
    gap: 12px;

    background-color: var(--background-300);
    border-radius: 24px;

    cursor: pointer;
    transition: 120ms ease-in-out;

    padding: 24px;

    &:hover {
      background-color: var(--background-200);
    }

    img {
      width: 48px;
      height: 48px;

      border-radius: 16px;

      user-select: none;
    }

    .info {
      display: flex;
      flex-direction: column;

      .name {
        display: flex;
        gap: 4px;

        h4 {
          color: var(--foreground-400);
        }
      }

      h4 {
        font: {
          size: 16px;
          weight: 600;
        };

        line-height: 24px;

        &.id {
          color: var(--foreground-100);
        }
      }
    }
  }
</style>