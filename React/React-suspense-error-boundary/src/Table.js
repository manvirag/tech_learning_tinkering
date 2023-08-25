import React from "react";
import { Table, Thead, Tbody, Tr, Th, Td, Box, Text } from "@chakra-ui/react";
import { SearchIcon } from "@chakra-ui/icons";
import { fetchData } from "./fakeApi";

const data = fetchData();

const EmptyState = () => {
  return (
    <Box textAlign="center">
      <SearchIcon fontSize="9xl" mt={20} color="blue.500" />
      <Text textAlign="center" my={10}>
        No cats found!
      </Text>
    </Box>
  );
};

const TableComponent = () => {
  const cats = data.cats.read();

  if (cats?.length === 0) {
    return <EmptyState />;
  }

  return (
    <Box p="3" w="100%" position="relative">
      <Table colorScheme="blue" overflow="none">
        <Thead>
          <Tr>
            <Th></Th>
            <Th>Name</Th>
            <Th>Origin</Th>
            <Th>Type</Th>
            <Th>Body type</Th>
            <Th>Coat</Th>
            <Th>Pattern</Th>
          </Tr>
        </Thead>
        <Tbody>
          {cats?.map((cat) => (
            <Tr key={cat.name}>
              <Td p="2">
                <Box
                  bgImage={cat.imgUrl}
                  w="150px"
                  h="150px"
                  backgroundSize="cover"
                />
              </Td>
              <Td>{cat.name}</Td>
              <Td>{cat.origin}</Td>
              <Td>{cat.type}</Td>
              <Td>{cat.bodyType}</Td>
              <Td>{cat.coat}</Td>
              <Td>{cat.pattern}</Td>
            </Tr>
          ))}
        </Tbody>
      </Table>
    </Box>
  );
};

export default TableComponent;
