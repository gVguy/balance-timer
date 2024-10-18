import { ref } from 'vue';
import { EventsOn } from '@w/runtime/runtime'

const session = ref({
  timeInitial: 0,
  timeRemaining: 0,
  timeRemainingString: '',
  state: 'Stopped',
})

EventsOn('time', (timeInt: number, timeString: string) => {
  console.log('Received "time" event', timeInt, timeString)
  session.value.timeRemaining = timeInt
  session.value.timeRemainingString = timeString
})

EventsOn('state', (state: string) => {
  console.log('Received "state" event', state)
  session.value.state = state
})

EventsOn('tickerStart', (timeInt: number, timeString: string) => {
  console.log('Received "tickerStart" event', timeInt, timeString)
  session.value.timeInitial = timeInt
  session.value.timeRemaining = timeInt
  session.value.timeRemainingString = timeString
})

export function useSession() {

  return {
    session
  }

}
