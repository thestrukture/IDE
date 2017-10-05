//
//  FlowProtocol.h
//  FlowCode
//
//  Created by Cheikh Seck on 4/2/15.
//  Copyright (c) 2015 Gopher Sauce LLC. All rights reserved.
//

#import <Foundation/Foundation.h>
#if TARGET_OS_IPHONE
#import <MobileCoreServices/MobileCoreServices.h>
#else
#import <CoreServices/CoreServices.h>
#endif
@interface FlowProtocol : NSURLProtocol

@property (readonly,copy) NSURLRequest *request;


@end
