<script lang="ts" setup>
import StoppedView from '@/components/views/StoppedView.vue'
import WorkingView from '@/components/views/WorkingView.vue'
import { useSession } from '@/composables/useSession'
import * as API from '@w/go/main/App'

;(window as any).api = API // expose go api methods for dev convenience

const { session } = useSession()

</script>

<template>
  <Transition name="view">
    <StoppedView v-if="session.state == 'Stopped'" />
    <WorkingView v-else-if="session.state == 'Working'" />
  </Transition>
</template>

<style scoped lang="scss">
.view-enter-from,
.view-leave-to {
  opacity: 0;
}
.view-leave-to {
  translate: 0 -50px;
}
.view-enter-from {
  translate: 0 50px;
}
.view-enter-active,
.view-leave-active {
  position: absolute;
  transition: all 1s, opacity .6s;
}
.view-enter-active {
  transition-delay: .3s;
}
</style>
