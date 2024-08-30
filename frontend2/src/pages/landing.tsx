import TaskCount from '../components/tasks/taskcount'
import TaskCount2 from '../components/tasks/taskcount2'


export default function LandingPage() {
    return (
        <main>
            <h1>Hello, Preact!</h1>
            <h2>Task count: <TaskCount/></h2>
            <h2>Task count 2.1: <TaskCount2/></h2>
            <h2>Task count 2.2: <TaskCount2/></h2>
        </main>
    )
  }
  