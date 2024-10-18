<script lang="ts" setup>
import {computed, onBeforeMount, ref} from 'vue'
import {GetConfigMinutes, StartSession} from '@w/go/main/App'
import { main } from '@w/go/models'
import ViewLayout from '@/components/ViewLayout.vue'

const config = ref<main.Config>()

onBeforeMount(async () => {
  config.value = await GetConfigMinutes()
  console.log('got config:', config.value)
})

function start() {
  console.log('start()')
  if (!config.value) return
  const { Working, Resting } = config.value.intervals
  StartSession(Working, Resting)
}

</script>

<template>
  <ViewLayout v-if="config">
    <template #heading>Setup a new session</template>
    <div class="block">
      <span class="label"><span class="highlight">Work</span> interval (min):</span>
      <input class="input" type="number" v-model="config.intervals.Working">
    </div>
    <div class="block">
      <span class="label"><span class="highlight">Rest</span> interval (min):</span>
      <input class="input" type="number" v-model="config.intervals.Resting">
    </div>
    <template #after>
      <button class="start-button" @click="start">Start</button>
    </template>
  </ViewLayout>
</template>

<style scoped lang="scss">
.block {
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  .label {
    color: rgba(var(--fg-rgb), .7);
    margin-bottom: 12px;
  }
}
.input {
  width: 92px;
  font-size: 18px;
}
</style>
