import React from "react";
import { MuiThemeProvider, createMuiTheme } from "@material-ui/core/styles";
import SignIn from "recipe-pages/home";
import "./App.css";

const theme = createMuiTheme();

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      <SignIn />
    </MuiThemeProvider>
  );
}

export default App;
