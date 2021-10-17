import React, { useState } from "react";
import { useHistory } from "react-router-dom";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import FormControlLabel from "@mui/material/FormControlLabel";
import Checkbox from "@mui/material/Checkbox";
import Link from "@mui/material/Link";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import { useAuth } from "../hooks/auth";
// import { DoorImage } from "../components/DoorImage";

// import "tailwindcss/dist/tailwind.css";

export function AuthLoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const auth = useAuth();
  const history = useHistory();

  function handleSubmit() {
    auth
      .login(username, password)
      .then(() => {
        history.replace("/todo");
        setUsername("");
        setPassword("");
      })
      .catch((err) => {
        console.log(err);
        // todo
      });
  }

  //   https://codepen.io/PortSpasy/pen/GRJyJyZ

  return (
    <Grid container component="main" sx={{ height: "100vh" }}>
      <Grid
        item
        xs={false}
        sm={4}
        md={7}
        sx={{
          backgroundImage: "url(https://source.unsplash.com/random)",
          backgroundRepeat: "no-repeat",
          backgroundColor: (t) => (t.palette.mode === "light" ? t.palette.grey[50] : t.palette.grey[900]),
          backgroundSize: "cover",
          backgroundPosition: "center",
        }}
      />
      <Grid item xs={12} sm={8} md={5} component={Paper} elevation={6} square>
        <Box
          sx={{
            my: 8,
            mx: 4,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box
            component="form"
            noValidate
            onSubmit={(e) => {
              e.preventDefault();
              handleSubmit();
            }}
            sx={{ mt: 1 }}
          >
            <TextField
              margin="normal"
              required
              fullWidth
              id="email"
              label="Email Address"
              name="email"
              autoComplete="email"
              autoFocus
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
            />
            <FormControlLabel control={<Checkbox value="remember" color="primary" />} label="Remember me" />
            <Button type="submit" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }}>
              Sign In
            </Button>
            <Grid container>
              <Grid item xs>
                <Link href="#" variant="body2">
                  Forgot password?
                </Link>
              </Grid>
              <Grid item>
                <Link href="#" variant="body2">
                  {"Don't have an account? Sign Up"}
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
      </Grid>
    </Grid>
  );

  // return (
  //   <div className="lg:flex">
  //     <div className="lg:w-1/2 xl:max-w-screen-sm">
  //       <div className="py-12 bg-indigo-100 lg:bg-white flex justify-center lg:justify-start lg:px-12">
  //         <div className="cursor-pointer flex items-center">
  //           <div>
  //             {/* <svg className="w-10 text-indigo-500" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" version="1.1" id="Layer_1" x="0px" y="0px" viewBox="0 0 225 225" style="enable-background:new 0 0 225 225;" xml:space="preserve">
  //                               <style type="text/css">
  //                                   .st0{fill:none;stroke:currentColor;stroke-width:20;stroke-linecap:round;stroke-miterlimit:3;}
  //                               </style>
  //                               <g transform="matrix( 1, 0, 0, 1, 0,0) ">
  //                               <g>
  //                               <path id="Layer0_0_1_STROKES" className="st0" d="M173.8,151.5l13.6-13.6 M35.4,89.9l29.1-29 M89.4,34.9v1 M137.4,187.9l-0.6-0.4     M36.6,138.7l0.2-0.2 M56.1,169.1l27.7-27.6 M63.8,111.5l74.3-74.4 M87.1,188.1L187.6,87.6 M110.8,114.5l57.8-57.8"/>
  //                               </g>
  //                               </g>
  //                           </svg> */}
  //           </div>
  //           <div className="text-2xl text-indigo-800 tracking-wide ml-2 font-semibold">blockify</div>
  //         </div>
  //       </div>
  //       <div className="mt-10 px-12 sm:px-24 md:px-48 lg:px-12 lg:mt-16 xl:px-24 xl:max-w-2xl">
  //         <h2
  //           className="text-center text-4xl text-indigo-900 font-display font-semibold lg:text-left xl:text-5xl
  //                   xl:text-bold"
  //         >
  //           Log in
  //         </h2>
  //         <div className="mt-12">
  //           <form>
  //             <div>
  //               <div className="text-sm font-bold text-gray-700 tracking-wide">Email Address</div>
  //               <input
  //                 className="w-full text-lg py-2 border-b border-gray-300 focus:outline-none focus:border-indigo-500"
  //                 type=""
  //                 placeholder="mike@example.com"
  //                 value={username}
  //                 onChange={(e) => setUsername(e.target.value)}
  //               />
  //             </div>
  //             <div className="mt-8">
  //               <div className="flex justify-between items-center">
  //                 <div className="text-sm font-bold text-gray-700 tracking-wide">Password</div>
  //                 <div>
  //                   <a
  //                     href="#TODO"
  //                     className="text-xs font-display font-semibold text-indigo-600 hover:text-indigo-800
  //                                       cursor-pointer"
  //                   >
  //                     Forgot Password?
  //                   </a>
  //                 </div>
  //               </div>
  //               <input
  //                 className="w-full text-lg py-2 border-b border-gray-300 focus:outline-none focus:border-indigo-500"
  //                 type=""
  //                 placeholder="Enter your password"
  //                 value={password}
  //                 onChange={(e) => setPassword(e.target.value)}
  //               />
  //             </div>
  //             <div className="mt-10">
  //               <button
  //                 className="bg-indigo-500 text-gray-100 p-4 w-full rounded-full tracking-wide
  //                               font-semibold font-display focus:outline-none focus:shadow-outline hover:bg-indigo-600
  //                               shadow-lg"
  //                 type="submit"
  //                 name="submit"
  //                 onClick={(e) => {
  //                   e.preventDefault();
  //                   handleLogin();
  //                 }}
  //               >
  //                 Log In
  //               </button>
  //             </div>
  //           </form>
  //           <div className="mt-12 text-sm font-display font-semibold text-gray-700 text-center">
  //             Don't have an account ?{" "}
  //             <a href="#TODO" className="cursor-pointer text-indigo-600 hover:text-indigo-800">
  //               Sign up
  //             </a>
  //           </div>
  //         </div>
  //       </div>
  //     </div>
  //     <div className="hidden lg:flex items-center justify-center bg-indigo-100 flex-1 h-screen">
  //       <DoorImage />
  //     </div>
  //   </div>
  // );
}
