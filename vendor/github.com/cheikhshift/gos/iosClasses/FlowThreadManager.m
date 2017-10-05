//
//  FlowThreadManager.m
//  FlowCode
//
//  Created by Cheikh Seck on 4/1/15.
//  Copyright (c) 2015 Gopher Sauce LLC. All rights reserved.
//

#import "FlowThreadManager.h"
#import <LocalAuthentication/LocalAuthentication.h>

@implementation FlowThreadManager {
   
}
static  NSMutableDictionary *dataset;
static  NSArray *threads;
static  FlowThreadManager *flowlayer;
static  FlowTissue *tissuelayer;
static Completion runAfter;

static CLGeocoder *geocoder;
static CLPlacemark *placemark;


+ (void) getGPS {
    //not working atm...
    FlowThreadManager *inj = [FlowThreadManager instance];
    if(inj.locationManager != nil){
        inj.locationManager = [[CLLocationManager alloc] init];
        geocoder = [[CLGeocoder alloc] init];
        inj.locationManager.delegate = inj;
        inj.locationManager.desiredAccuracy = kCLLocationAccuracyBest;
        inj.locationManager.distanceFilter = kCLDistanceFilterNone;
        [inj.locationManager requestWhenInUseAuthorization];
        [inj.locationManager startUpdatingLocation];
    }
    
}

- (void)accelerometer:(UIAccelerometer *)accelerometer didAccelerate:
(UIAcceleration *)acceleration{
        /* call endpoint update */
        //Will use js Subset to handle GPS and accelerometer
}

#pragma mark - CLLocationManagerDelegate

- (void)locationManager:(CLLocationManager *)manager didFailWithError:(NSError *)error
{
    NSLog(@"GPS didFailWithError: %@", error);
}

- (void)locationManager:(CLLocationManager *)manager didUpdateToLocation:(CLLocation *)newLocation fromLocation:(CLLocation *)oldLocation
{
    NSLog(@"didUpdateToLocation: %@", newLocation);
    CLLocation *currentLocation = newLocation;
    
    if (currentLocation != nil) {
      //  longitudeLabel.text = [NSString stringWithFormat:@"%.8f", currentLocation.coordinate.longitude];
       // latitudeLabel.text = [NSString stringWithFormat:@"%.8f", currentLocation.coordinate.latitude];
    }
    /*
     
     addressLabel.text = [NSString stringWithFormat:@"%@ %@\n%@ %@\n%@\n%@",
     placemark.subThoroughfare, placemark.thoroughfare,
     placemark.postalCode, placemark.locality,
     placemark.administrativeArea,
     placemark.country];
     */
    // Stop Location Manager
    [self.locationManager stopUpdatingLocation];
    
    NSLog(@"Resolving the Address");
    [geocoder reverseGeocodeLocation:currentLocation completionHandler:^(NSArray *placemarks, NSError *error) {
        NSLog(@"Found placemarks: %@, error: %@", placemarks, error);
        if (error == nil && [placemarks count] > 0) {
            placemark = [placemarks lastObject];
            //callback via Go end point...
            } else {
            NSLog(@"%@", error.debugDescription);
        }
    } ];
    
}

+ (void) createFlowLayer {
    if(![flowlayer isKindOfClass:[FlowThreadManager class]]){
        flowlayer = [[FlowThreadManager alloc] init];
    }
}

+ (void) takePicture:(NSString *)name {
    [FlowThreadManager instance].tempstring = name;
    
    if (![UIImagePickerController isSourceTypeAvailable:UIImagePickerControllerSourceTypeCamera]) {
        
        UIAlertView *myAlertView = [[UIAlertView alloc] initWithTitle:@"Error"
                                                              message:@"Device has no camera"
                                                             delegate:nil
                                                    cancelButtonTitle:@"OK"
                                                    otherButtonTitles: nil];
        
        [myAlertView show];
        
    }
    else {
        UIImagePickerController *picker = [[UIImagePickerController alloc] init];
        picker.delegate = [FlowThreadManager currentFlowEnclosing];
        picker.allowsEditing = YES;
       
        picker.sourceType = UIImagePickerControllerSourceTypeCamera;
        
        [[FlowThreadManager currentFlowEnclosing] presentViewController:picker animated:YES completion:NULL];
    }
}

+ (ViewController*) currentFlowEnclosing {
    UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
    // Replace the current view controller
    //viewdidappear
    NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
    
    ViewController *viewrs;
    if([[viewControllers lastObject] isKindOfClass:[ViewController class]]){
        viewrs = [viewControllers lastObject];
    }
    else  viewrs = [viewControllers objectAtIndex:[viewControllers count] - 2];
    return viewrs;
}

+ (id) tissue {
    if(![tissuelayer isKindOfClass:[FlowTissue class]]){
        tissuelayer = [[FlowTissue alloc] init];
    }
    return tissuelayer;
}

+(FlowThreadManager *) instance {
    [self createFlowLayer];
    return flowlayer;
}

+ (id) getobject:(NSString *) name {
    return dataset[name];
}

+ (BOOL) saveobject:(id)object withName:(NSString *) key {
    //key
    dataset[key] = object;
    return YES;
}

-(void) webViewDidStartLoad:(UIWebView *)webView {
    //proccesspage
}
- (BOOL)webView:(UIWebView *)webView shouldStartLoadWithRequest:(NSURLRequest *)request navigationType:(UIWebViewNavigationType)navigationType {
    
    NSURL *url = [request URL];
    NSString *urlStr = url.absoluteString;
    
    return [self processURL:urlStr];
    
}

+ (void) runJSBackone:(NSString *) js {
    
    UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
    // Replace the current view controller
    //viewdidappear
    NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
    ViewController *viewrs = [viewControllers objectAtIndex:[viewControllers count] - 2];
    [viewrs.webView stringByEvaluatingJavaScriptFromString:js];
}


+ (NSString*)makeParamtersString:(NSDictionary*)parameters withEncoding:(NSStringEncoding)encoding
{
    if (nil == parameters || [parameters count] == 0)
        return nil;
    
    NSMutableString* stringOfParamters = [[NSMutableString alloc] init];
    NSEnumerator *keyEnumerator = [parameters keyEnumerator];
    id key = nil;
    while ((key = [keyEnumerator nextObject]))
    {
        NSString *value = [[parameters valueForKey:key] isKindOfClass:[NSString class]] ?
        [parameters valueForKey:key] : [[parameters valueForKey:key] stringValue];
        [stringOfParamters appendFormat:@"%@=%@&",
         [self URLEscaped:key withEncoding:encoding],
         [self URLEscaped:value withEncoding:encoding]];
    }
    
    // Delete last character of '&'
    NSRange lastCharRange = {[stringOfParamters length] - 1, 1};
    [stringOfParamters deleteCharactersInRange:lastCharRange];
    return stringOfParamters;
}

+ (NSString *)URLEscaped:(NSString *)strIn withEncoding:(NSStringEncoding)encoding
{
    CFStringRef escaped = CFURLCreateStringByAddingPercentEscapes(NULL, (CFStringRef)strIn, NULL, (CFStringRef)@"!*'();:@&=+$,/?%#[]", CFStringConvertNSStringEncodingToEncoding(encoding));
    NSString *strOut = [NSString stringWithString:(__bridge NSString *)escaped];
    CFRelease(escaped);
    return strOut;
}

- (BOOL) processURL:(NSString *) url
{
    NSString *urlStr = [NSString stringWithString:url];
    
    NSString *protocolPrefix = @"flowcode://";
    
    //process only our custom protocol
    if ([[urlStr lowercaseString] hasPrefix:protocolPrefix])
    {
        //strip protocol from the URL. We will get input to call a native method
        urlStr = [urlStr substringFromIndex:protocolPrefix.length];
        
        //Decode the url string
        urlStr = [urlStr stringByReplacingPercentEscapesUsingEncoding:NSUTF8StringEncoding];
        
        NSError *jsonError;
        
        //parse JSON input in the URL
        NSDictionary *callInfo = [NSJSONSerialization
                                  JSONObjectWithData:[urlStr dataUsingEncoding:NSUTF8StringEncoding]
                                  options:kNilOptions
                                  error:&jsonError];
        
        //check if there was error in parsing JSON input
        if (jsonError != nil)
        {
            NSLog(@"Error parsing JSON for the url %@",url);
            return NO;
        }
        
        //Get function name. It is a required input
        NSString *functionName = [callInfo objectForKey:@"functionname"];
        if (functionName == nil)
        {
            NSLog(@"Missing function name");
            return NO;
        }
        
        NSString *successCallback = [callInfo objectForKey:@"success"];
        NSString *errorCallback = [callInfo objectForKey:@"error"];
        NSArray *argsArray = [callInfo objectForKey:@"args"];
        
        [self callNativeFunction:functionName withArgs:argsArray onSuccess:successCallback onError:errorCallback];
        
        //Do not load this url in the WebView
        return NO;
        
    }
    
    return YES;
}


- (void)webView:(UIWebView *)webView
didFailLoadWithError:(NSError *)error {
    NSLog(@"this happened %@", error);
}

-(void) webViewDidFinishLoad:(UIWebView *)webView {
  //  NSLog(@"rrr!");
   
    if(runAfter){
    runAfter();
    runAfter = nil;
    }
}

+ (void) webviewCompletion:(Completion) finished {
    runAfter = ^(void){
        finished();
    };
}

+ (NSString *) flowspin:(NSString *) view {
    NSString *work = view;
    //process javascript here @{}; @+;
    
    return work;
}


- (void)connection:(NSURLConnection *)connection didFailWithError:(NSError *)error {
    // The request has failed for some reason!
    // Check the error var
}

- (void)paymentMethodCreator:(id)sender requestsPresentationOfViewController:(UIViewController *)viewController {
    UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
    // Replace the current view controller
    //viewdidappear
    NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
    ViewController *viewrs = [viewControllers lastObject];
    [viewrs presentViewController:viewController animated:YES completion:nil];
}

+ (ViewController *) upperStack {
    
    UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
    // Replace the current view controller
    //viewdidappear
    NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
    ViewController *viewrs = [viewControllers lastObject];
    return viewrs;
    
}


- (void)paymentMethodCreator:(id)sender requestsDismissalOfViewController:(UIViewController *)viewController {
    [viewController dismissViewControllerAnimated:YES completion:nil];
}

+ (void) pulseView : (NSString *) url {
    UIStoryboard *st = [UIStoryboard storyboardWithName:@"Main"
                                                 bundle:[NSBundle mainBundle]];
    
    ViewController *viewController =  [st instantiateViewControllerWithIdentifier:@"FlowView"];
    viewController.override = YES;
    viewController.viewurl = [NSString stringWithFormat:@"http://localhost/%@", url];
    [(UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController pushViewController:viewController animated:YES];
}




+ (NSString *) flowjs:(NSString *) url withData:(NSArray *) args {
    NSString *initialjs = [NSString stringWithFormat:@"var form = %@",args];
    NSDictionary *payload = (NSDictionary*) args;
    if([payload objectForKey:@"url"] && [url isEqualToString:@"flush"]){
        
        NSData *jsonData =[NSJSONSerialization dataWithJSONObject:payload[@"data"] options:0 error:nil];
        
        NSString *jsonStr1 = [NSString stringWithUTF8String:[jsonData bytes]];
        
        NSString *imagePath = [[[[NSBundle mainBundle] bundlePath] stringByAppendingString:@"/SharedCode/Views"] stringByAppendingString:[payload objectForKey:@"url"] ];
        NSData *data = [NSData dataWithContentsOfFile:imagePath];
       // NSLog(@"Pong %@ ", jsonStr1);
        initialjs = [[NSString  stringWithFormat:@"var form = JSON.parse('%@');",jsonStr1] stringByAppendingString:[[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding]];
        //get htmlstring and append
    }
    else if([payload objectForKey:@"targ"] && [url isEqualToString:@"lajax"]){
        
              NSString *imagePath = [[[[NSBundle mainBundle] bundlePath] stringByAppendingString:@"/SharedCode/Views"] stringByAppendingString:[payload objectForKey:@"targ"] ];
        NSData *data = [NSData dataWithContentsOfFile:imagePath];
        initialjs =  [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding];
        //get htmlstring and append
    }
    else if([payload objectForKey:@"name"] && [url isEqualToString:@"nssave"]){
        NSData *jsonData =[NSJSONSerialization dataWithJSONObject:payload[@"payload"] options:0 error:nil];
        
        
        [[NSUserDefaults standardUserDefaults] setObject:jsonData forKey:[payload objectForKey:@"name"]];
       
        [[NSUserDefaults standardUserDefaults] synchronize];
        //get htmlstring and append
    }
    else if([url isEqualToString:@"fingerscan"]){
        //345600
        NSUserDefaults *flowsc = [NSUserDefaults standardUserDefaults];

        double lastdate = [flowsc doubleForKey:@"fingerGate"];
        
        if((round([[NSDate date] timeIntervalSince1970]) - lastdate ) > 172800  ){
            
   
        LAContext *context = [[LAContext alloc] init];
        NSError *error = nil;
        
        if ([context canEvaluatePolicy:LAPolicyDeviceOwnerAuthenticationWithBiometrics error:&error]) {
            // Authenticate User
            
            [context evaluatePolicy:LAPolicyDeviceOwnerAuthenticationWithBiometrics localizedReason:@"Please authenticate to view your account" reply:^(BOOL success, NSError *authenticationError){
                
                dispatch_async(dispatch_get_main_queue(), ^{
                //suppress if sign last signin is less then x seconds
                if (success) {
                     [[NSUserDefaults standardUserDefaults] setDouble:round([[NSDate date] timeIntervalSince1970]) forKey:@"fingerGate"];
                     [FlowThreadManager runJS:@"fingerwork()"];
                }
                else {
                  //  NSLog(@"Fingerprint validation failed: %@.", authenticationError.localizedDescription);
                    if(authenticationError.code == LAErrorUserFallback){
                        //open alert and based on the value is there
                        //two alerts with fields
                        [FlowThreadManager pinLogin];
                    }
                    else
                    [FlowThreadManager runJS:@"fingerfail()"];
                    
                }
                
                });
            }];
            
        } else {
           //no message
           //go to pin screen // pin.html
            //
            [FlowThreadManager pinLogin];
            
        }
        
        }
        else {
           // NSLog(@"ohh no %f", lastdate);
        }
    }
    else if([payload objectForKey:@"url"] && [url isEqualToString:@"fajax"]){
        // Create the request.
        NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:[NSURL URLWithString:payload[@"url"]]];
        
        // Specify that it will be a POST request
        request.HTTPMethod = @"POST";
        
        // This is how we set header fields
        [request setValue:@"application/xml; charset=utf-8" forHTTPHeaderField:@"Content-Type"];
        
        // Convert your data and set your request's HTTPBody property
        NSString *stringData = [self makeParamtersString:payload[@"payload"] withEncoding:NSUTF8StringEncoding];;
        NSData *requestBodyData = [stringData dataUsingEncoding:NSUTF8StringEncoding];
        request.HTTPBody = requestBodyData;
        
        // Create url connection and fire request
        /*
         NSLog(@"success");
         [FlowThreadManager runJS:[NSString stringWithFormat:@"succF('%@')", [[NSString alloc] initWithData:_responseData encoding:NSUTF8StringEncoding] ] ];
         */
      
        
        dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{
            //proccess string get html
            //evaluate Flowjs
            NSURLResponse * response = nil;
            NSError * error = nil;
                NSData * data = [NSURLConnection sendSynchronousRequest:request
                                              returningResponse:&response
                                                          error:&error];
            NSLog(@"%@", error);
             NSLog(@"resp %@", response);
        if (error == nil)
        {
            // Parse data here
            NSLog(@"success");
            [FlowThreadManager runJS:[NSString stringWithFormat:@"succF('%@')", [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding] ] ];
        }
            
        });
        //get htmlstring and append
    }
    else if([url isEqualToString:@"safariopen"]){
        [[UIApplication sharedApplication] openURL:[NSURL URLWithString:payload[@"url"]] ];
    }
    else if([url isEqualToString:@"ParseCI"]){
        //save here
     
    }
    else if([payload objectForKey:@"name"] && [url isEqualToString:@"nsopen"]){
        if([[NSUserDefaults standardUserDefaults] objectForKey:[payload objectForKey:@"name"]]){
            NSString *jsonStr1 = [NSString stringWithUTF8String:[[[NSUserDefaults standardUserDefaults] objectForKey:[payload objectForKey:@"name"]] bytes] ];
           initialjs = jsonStr1;
  //   NSLog(@"data watch %@", initialjs);
        }
        else initialjs = @"{}";
        //get htmlstring and append
       //   NSLog(@"data watch %@", [[NSUserDefaults standardUserDefaults] objectForKey:[payload objectForKey:@"name"]]);
    }
   /* else if([url isEqualToString:@"showload"]){
        [SwiftSpinner show:[payload objectForKey:@"text"] animated:YES];
    }
    else if([url isEqualToString:@"showloadstatic"]){
        [SwiftSpinner show:[payload objectForKey:@"text"] animated:NO];
    }
    else if([url isEqualToString:@"hideload"]){
        [SwiftSpinner hide];
    } */
    else if([url isEqualToString:@"dismiss"]){
        UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
        // Replace the current view controller
        NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
      
        [viewControllers removeLastObject];
     
        [navcontroller setViewControllers:viewControllers animated:YES];
        
    }
    else if([url isEqualToString:@"clearnotifs"]){
        [[NSUserDefaults standardUserDefaults] setObject:@"NO" forKey:@"notif"];
        [[NSUserDefaults standardUserDefaults] synchronize];
    }
    else if([url isEqualToString:@"rootview"]){
        UIStoryboard *st = [UIStoryboard storyboardWithName:@"Main"
                                                     bundle:[NSBundle mainBundle]];
        
        ViewController *viewController =  [st instantiateViewControllerWithIdentifier:@"FlowView"];
        viewController.override = YES;
        viewController.viewurl = payload[@"url"];
        [(UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController pushViewController:viewController animated:YES];
        initialjs = @"";
        /*
         UIStoryboard *sb = [UIStoryboard storyboardWithName:@"NameOfStoryBoard" bundle:nil];
         UITabBarController *rootViewController = [sb instantiateViewControllerWithIdentifier:@"NameOfTabBarController"];
         [[UIApplication sharedApplication].keyWindow setRootViewController:rootViewController];
         */
    }
    
     //NSLog(@"%@",initialjs);
    return initialjs;
}

+ (void) pinLogin {
    NSString *text;
    if([[NSUserDefaults standardUserDefaults] objectForKey:@"impin"])
        text = @"Please enter the pin you saved on this application to unlock it";
    else text = @"No current pin is set. Please enter a pin number for enhanced security on orenda";
    UIAlertView * alert =[[UIAlertView alloc ] initWithTitle:@"Pin Login" message:text delegate:self cancelButtonTitle:@"Validate" otherButtonTitles: nil];
    alert.alertViewStyle = UIAlertViewStyleSecureTextInput;
    //[alert addButtonWithTitle:@"Login"];
    UITextField *username = [alert textFieldAtIndex:0];
    username.keyboardType = UIKeyboardTypeNumberPad;
    username.placeholder = @"Pin";
    username.textAlignment = 1;
    [alert show];
}
+ (void) WrongpinLogin {
    NSString *text = @"Incorrect pin entered, please try again!";
      UIAlertView * alert =[[UIAlertView alloc ] initWithTitle:@"Pin Login" message:text delegate:self cancelButtonTitle:@"Validate" otherButtonTitles: nil];
    alert.alertViewStyle = UIAlertViewStyleSecureTextInput;
    //[alert addButtonWithTitle:@"Login"];
    UITextField *username = [alert textFieldAtIndex:0];
    username.keyboardType = UIKeyboardTypeNumberPad;
    username.placeholder = @"Pin";
    username.textAlignment = 1;
    [alert show];
}

+ (void)alertView:(UIAlertView *)alertView didDismissWithButtonIndex:(NSInteger)buttonIndex
{
    //process
    UITextField *username = [alertView textFieldAtIndex:0];
    if([username.text isEqualToString:@""]){
        [FlowThreadManager pinLogin];
    }
    else {
        //pin there
        if([[NSUserDefaults standardUserDefaults] objectForKey:@"impin"]){
            //compare
            if(![[[NSUserDefaults standardUserDefaults] objectForKey:@"impin"] isEqualToString:username.text]){
                [FlowThreadManager WrongpinLogin];
            }
            else {
                //worked
            [[NSUserDefaults standardUserDefaults] setDouble:round([[NSDate date] timeIntervalSince1970]) forKey:@"fingerGate"];
            }
        }
        else {
           //save pin and auth
            [[NSUserDefaults standardUserDefaults] setObject:username.text forKey:@"impin"];
                                 [[NSUserDefaults standardUserDefaults] setDouble:round([[NSDate date] timeIntervalSince1970]) forKey:@"fingerGate"];
        }
    }
   
}

- (void) callNativeFunction:(NSString *) name withArgs:(NSArray *) args onSuccess:(NSString *) successCallback onError:(NSString *) errorCallback
{
    //We only know how to process sayHello
    
     [self callSuccessCallback:successCallback withRetValue:[FlowThreadManager flowjs:name withData:args] forFunction:name];
         
    
    /* if ([name compare:@"sayHello" options:NSCaseInsensitiveSearch] == NSOrderedSame)
    {
        if (args.count > 0)
        {
            NSString *resultStr = [NSString stringWithFormat:@"Hello %@ !", [args objectAtIndex:0]];
            
            [self callSuccessCallback:successCallback withRetValue:resultStr forFunction:name];
        }
        else
        {
            NSString *resultStr = [NSString stringWithFormat:@"Error calling function %@. Error : Missing argument", name];
            [self callErrorCallback:errorCallback withMessage:resultStr];
        }
    }
    else
    {
        //Unknown function called from JavaScript
        NSString *resultStr = [NSString stringWithFormat:@"Cannot process function %@. Function not found", name];
        [self callErrorCallback:errorCallback withMessage:resultStr];
        
    } */
}

-(void) callErrorCallback:(NSString *) name withMessage:(NSString *) msg
{
    if (name != nil)
    {
        //call error handler
        
        NSMutableDictionary *resultDict = [[NSMutableDictionary alloc] init];
        [resultDict setObject:msg forKey:@"error"];
        [self callJSFunction:name withArgs:resultDict];
    }
    else
    {
        NSLog(@"%@",msg);
    }
    
}

-(void) callSuccessCallback:(NSString *) name withRetValue:(id) retValue forFunction:(NSString *) funcName
{
    if (name != nil)
    {
        //call succes handler
        
        NSMutableDictionary *resultDict = [[NSMutableDictionary alloc] init];
        [resultDict setObject:retValue forKey:@"result"];
        [self callJSFunction:name withArgs:resultDict];
    }
    else
    {
       // NSLog(@"Result of function %@ = %@", funcName,retValue);
    }
    
}

+ (void) loadScreen:(BOOL)switc usingMessage:(NSString *)message {
    if(switc){
        [DejalBezelActivityView activityViewForView:[FlowThreadManager currentFlowEnclosing].view];
    }
    else {
        [DejalBezelActivityView removeView];
    }
}

+ (void) runJS:(NSString *)function {
    [[FlowThreadManager currentFlow] stringByEvaluatingJavaScriptFromString:function];
}



-(void) callJSFunction:(NSString *) name withArgs:(NSMutableDictionary *) args
{
    NSError *jsonError;
    
    NSData *jsonData = [NSJSONSerialization dataWithJSONObject:args options:0 error:&jsonError];
    
    if (jsonError != nil)
    {
        //call error callback function here
        NSLog(@"Error creating JSON from the response  : %@",[jsonError localizedDescription]);
        return;
    }
    
    //initWithBytes:length:encoding
    NSString *jsonStr = [[NSString alloc] initWithData:jsonData encoding:NSUTF8StringEncoding];
    
   // NSLog(@"jsonStr = %@", json);

    if (jsonStr == nil)
    {
        //NSLog(@"jsonStr is null. count = %d", [args count]);
    }
    else {
        UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
        if([navcontroller isKindOfClass:[UINavigationController class]]){
        // Replace the current view controller
        NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
        ViewController *viewrs = [viewControllers lastObject];
    [viewrs.webView stringByEvaluatingJavaScriptFromString:[NSString stringWithFormat:@"%@(\"%@\");",name,[self escapeStringForJavascript:jsonStr] ]];
        }
    }
}





+ (UIWebView *) currentFlow {
    UINavigationController *navcontroller = (UINavigationController *)[UIApplication sharedApplication].keyWindow.rootViewController;
    // Replace the current view controller
    //viewdidappear
    NSMutableArray *viewControllers = [NSMutableArray arrayWithArray:[navcontroller viewControllers]];
    
    ViewController *viewrs;
    if([[viewControllers lastObject] isKindOfClass:[ViewController class]]){
    viewrs = [viewControllers lastObject];
    }
    else  viewrs = [viewControllers objectAtIndex:[viewControllers count] - 2];
    return viewrs.webView;
}

- (NSData*) encryptString:(NSString*)plaintext withKey:(NSString*)key {
    return [[plaintext dataUsingEncoding:NSUTF8StringEncoding] base64EncodedDataWithOptions:0];
}

- (NSString*) escapeStringForJavascript:(NSString*)input
{
    NSMutableString* ret = [NSMutableString string];
    int i;
    for (i = 0; i < input.length; i++)
    {
        unichar c = [input characterAtIndex:i];
        if (c == '\\')
        {
            // escape backslash
            [ret appendFormat:@"\\\\"];
        }
        else if (c == '"')
        {
            // escape double quotes
            [ret appendFormat:@"\\\""];
        }
        else if (c >= 0x20 && c <= 0x7E)
        {
            // if it is a printable ASCII character (other than \ and "), append directly
            [ret appendFormat:@"%c", c];
        }
        else
        {
            // if it is not printable standard ASCII, append as a unicode escape sequence
            [ret appendFormat:@"\\u%04X", c];
        }
    }
    return ret;
}



+ (void) process: (NSString *) mvc completion:(CodeProcessCompletion) finished {
    //dispatch async and run
    
    dispatch_async(dispatch_get_global_queue(DISPATCH_QUEUE_PRIORITY_DEFAULT, 0), ^{
        //proccess string get html
        //evaluate Flowjs
        NSString *result = [self flowspin:mvc];
        finished(result);
        //load string now on webview
    });

}




//.m





@end
