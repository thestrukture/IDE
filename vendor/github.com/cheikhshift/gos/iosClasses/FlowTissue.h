
			//
			//  FlowTissue.h
			//  GoTetst2
			//
			//  Created by OrendaCapital on 12/29/15.
			//  Copyright © 2015 Cheikh Seck LLC. All rights reserved.
			//

			#import <Foundation/Foundation.h>
			#import <AVFoundation/AVFoundation.h>
			#import <CoreLocation/CoreLocation.h>
			#import "Mymobile/Mymobile.h"
			#import "ViewController.h"
			#import "FlowThreadManager.h"


			@interface FlowTissue : NSObject  <GoMymobileFlow> {
			    
			}

			+ (void) handleRequest:(NSString *) endpoint;
			@end

	