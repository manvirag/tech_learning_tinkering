import {
  compose,
  defaultProps,
  mapProps,
  withHandlers,
  withState
} from "recompose";

//https://github.com/acdlite/recompose/blob/master/docs/API.md
// order left to right , don't know why

/*

will get the props after left side component and will send it 
to right side component after map

*/

const getConditions = mapProps((props) => {
  console.log("in map props");
  console.log(props); // will contain first as well
  return {
    ...props,
    firstMap: "hello"
  };
});

/*
what is we have two mapprops what will be their order
*/

const getConditions2 = mapProps((props) => {
  console.log("in second map props");
  console.log(props); // will contain first as well
  return {
    ...props,
    secondMap: "hello_2"
  };
});

/*
add defined props in left side component prop
and then send it to right side
*/
const addDefaultType = defaultProps({
  profileType: "PROFILE_TYPES.TEXT"
});

/*
withState(
  stateName: string,
  stateUpdaterName: string,
  initialState: any | (props: Object) => any
): HigherOrderComponent


add react state and its updater in  props of left side component prop
and then send it to right side


*/
const addState = withState("valueMap", "updateValueMap", (props) => {
  console.log("in add state");
  console.log(props);

  return "valueMapKiValue";
});

/*

withHandlers(
  handlerCreators: {
    [handlerName: string]: (props: Object) => Function
  } |
  handlerCreatorsFactory: (initialProps) => {
    [handlerName: string]: (props: Object) => Function
  }
): HigherOrderComponent



Same take the prop of left side component add handler 
and with use of their prop, then send handler + left component
prop to right
*/
const addHandlers = withHandlers({
  onConditionTypeChange: (props) => (e) => {
    console.log("in add handler");
    console.log(props);
    props.onChange(e);
  }
});

const withFirstProp = (Child) => {
  const WithFirstProp = (props) => {
    console.log("WithFirstProp");
    console.log(props);
    return <Child {...props} first={"first"} />;
  };
  return WithFirstProp;
};

const App = (props) => {
  console.log("see what props are coming");
  console.log(props);
  return (
    <>
      <h1>Hello CodeSandbox</h1>
      <h2>Start editing to see some magic happen!</h2>
    </>
  );
};

export default compose(
  withFirstProp,
  addHandlers,
  addState,
  addDefaultType,
  getConditions2,
  getConditions
)(App);
