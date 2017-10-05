//
//  ViewController.h
//  Go Mobile Test
//
//  Created by OrendaCapital on 12/13/15.
//  Copyright © 2015 Cheikh Seck LLC. All rights reserved.
//

#import <UIKit/UIKit.h>
#import "FlowThreadManager.h"
#import "FlowProtocol.h"


@interface ViewController : UIViewController  <UIImagePickerControllerDelegate, UINavigationControllerDelegate>

@property (weak, nonatomic) IBOutlet UIWebView *webView;
@property BOOL override;
@property NSString * viewurl;

@end

