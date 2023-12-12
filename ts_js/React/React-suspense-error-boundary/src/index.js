import React, { Suspense } from "react";
import {
  ChakraProvider,
  Center,
  Spinner,
  Alert,
  AlertIcon,
  AlertTitle,
  AlertDescription
} from "@chakra-ui/react";
import * as ReactDOM from "react-dom";
import "./styles.css";
import Table from "./Table";
import ErrorBoundary from "./ErrorBoundary";

const LoadingState = () => (
  <Center h="100vh">
    <Spinner margin="auto" emptyColor="gray.200" color="blue.500" size="xl" />
  </Center>
);

const ErrorState = () => {
  return (
    <Alert status="error">
      <AlertIcon />
      <AlertTitle mr={2}>An error occured!</AlertTitle>
      <AlertDescription>Please contact us for assistance.</AlertDescription>
    </Alert>
  );
};

const rootElement = document.getElementById("root");

ReactDOM.createRoot(rootElement).render(
  <ChakraProvider>
    <ErrorBoundary fallback={<ErrorState />}>
      <Suspense fallback={<LoadingState />}>
        <Table />
      </Suspense>
    </ErrorBoundary>
  </ChakraProvider>
);
