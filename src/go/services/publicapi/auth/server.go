package auth

// https://medium.com/@amsokol.com/tutorial-part-3-how-to-develop-go-grpc-microservice-with-http-rest-endpoint-middleware-739aac8f1d7e

// func Server() {
// 	logger.InitZapGlobal(logger.LevelDebug, time.RFC3339Nano)
// 	runServer()
// }

// func connectUserService() (core.UserServiceClient, *grpc.ClientConn) {
// 	opts := []grpc.DialOption{
// 		grpc.WithInsecure(),
// 	}

// 	host := util.Getenv("USER_SERVICE_ADDR", ":13001")
// 	conn, err := grpc.Dial(host, opts...)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	svc := core.NewUserServiceClient(conn)

// 	return svc, conn
// }

// func connectOAuthService() (core.OauthUserServiceClient, *grpc.ClientConn) {
// 	opts := []grpc.DialOption{
// 		grpc.WithInsecure(),
// 	}

// 	host := util.Getenv("OAUTH_USER_SERVICE_ADDR", ":13002")
// 	conn, err := grpc.Dial(host, opts...)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	svc := core.NewOauthUserServiceClient(conn)

// 	return svc, conn
// }

// func runServer() {
// 	ctx := context.Background()
// 	port := util.Getenv("PORT", "14586")
// 	lis, err := net.Listen("tcp", ":"+port)
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}

// 	// Connect to the user service
// 	userService, userConn := connectUserService()
// 	defer userConn.Close()
// 	oauthService, oauthConn := connectOAuthService()
// 	defer oauthConn.Close()

// 	// Connect to redis
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     util.Getenv("REDIS_ADDR", "localhost:6379"),
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})

// 	s := NewAuthenticationServer(userService, oauthService, NewAttemptCounter("publicapi:authentication", rdb))

// 	// gRPC server statup options
// 	opts := []grpc.ServerOption{}

// 	// add middleware
// 	opts = middleware.AddLogging(logger.ZapLogger, opts)

// 	grpcServer := grpc.NewServer(opts...)

// 	// attach the Todo service
// 	publicapi.RegisterAuthenticationServiceServer(grpcServer, &s)
// 	grpc_health_v1.RegisterHealthServer(grpcServer, grpcutil.NewServer())

// 	// graceful shutdown
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	go func() {
// 		for range c {
// 			// sig is a ^C, handle it
// 			logger.ZapLogger.Info("shutting down gRPC server...")

// 			grpcServer.GracefulStop()

// 			<-ctx.Done()
// 		}
// 	}()

// 	logger.ZapLogger.Info("staring gRPC server... port=" + port)

// 	if err := grpcServer.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %s", err)
// 	} else {
// 		log.Printf("Server started successfully")
// 	}
// }
