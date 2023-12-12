import "./styles.css";
import LearnCompose from "./LearnCompose";

export default function App() {
  return (
    <div className="App">
      <LearnCompose
        parent={"parent"}
        onChange={(data) => console.log("I am in parent " + data)}
      />
    </div>
  );
}
