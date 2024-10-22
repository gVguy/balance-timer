<script lang="ts" setup>
import {computed, onBeforeMount, ref, watchEffect} from 'vue'
import {GetConfigMinutes, StartSession} from '@w/go/main/App'
import ViewLayout from '@/components/ViewLayout.vue'
import { useSession } from '@/composables/useSession'

const { session } = useSession()

onBeforeMount(async () => {
  const configMinutes = await GetConfigMinutes()
  session.value.config.workMin = configMinutes.intervals.Working
  session.value.config.restMin = configMinutes.intervals.Resting
  console.log('got config (min):', configMinutes)
})

function start() {
  console.log('start()')
  const { workMin, restMin } = session.value.config
  StartSession(workMin, restMin)
}

const validate = (value: number) => {
  return Number.isInteger(value) && value > 0
}

const isValidWork = computed(() => validate(session.value.config.workMin))
const isValidRest = computed(() => validate(session.value.config.restMin))
const isValidAll = computed(() => isValidWork.value && isValidRest.value)

watchEffect(() => {
  console.log('isValidWork', isValidWork.value, 'session.value.config.workMin', session.value.config.workMin)
})

</script>

<template>
  <ViewLayout>
    <template #heading>Start a new session</template>
    <div class="block">
      <span class="label"><span class="highlight">Work</span> interval (min):</span>
      <input
        class="input"
        type="number"
        v-model="session.config.workMin"
        :class="{ 'not-valid': !isValidWork }"
      >
    </div>
    <div class="block">
      <span class="label"><span class="highlight">Rest</span> interval (min):</span>
      <input
        class="input"
        type="number"
        v-model="session.config.restMin"
        :class="{ 'not-valid': !isValidRest }"
      >
    </div>
    <template #actions>
      <button
        class="start-button"
        :class="{ 'inactive': !isValidAll }"
        @click="start"
      >
        Start
      </button>
    </template>
    <template #afterActions>
      <span class="hint" :class="{ visible: !isValidAll }">
        Please use positive number values
      </span>
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
  &.not-valid {
    border-color: indianred;
  }
}
.start-button {
  margin-top: 12px;
}
.hint {
  opacity: 0;
  transition: opacity .4s;
  color: indianred;
  height: 0;
  &.visible {
    opacity: 1;
  }
}
</style>
