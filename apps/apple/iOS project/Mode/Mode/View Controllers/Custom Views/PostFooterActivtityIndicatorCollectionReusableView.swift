//
//  PostFooterActivtityIndicatorCollectionReusableView.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/15/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class PostFooterActivtityIndicatorCollectionReusableView: UICollectionReusableView {
    
    // MARK: - Outlets
    
    @IBOutlet weak var loadMorePostsActivityIndicator: UIActivityIndicatorView!
    
    // MARK: - Properties
    
//    var isAnimatingFinal: Bool = false
//    var currentTransform: CGAffineTransform?
    
    // MARK: - Lifecycle
    
    override func awakeFromNib() {
        super.awakeFromNib()
        loadMorePostsActivityIndicator.startAnimating()
    }
//
//    override func layoutSubviews() {
//        super.layoutSubviews()
//    }
//
//    // MARK: - Custom Functions
//
//    func setTransform(inTransform: CGAffineTransform, scaleFactor: CGFloat) {
//        if isAnimatingFinal {
//            return
//        }
//        currentTransform = inTransform
//        loadMorePostsActivityIndicator.transform = CGAffineTransform(scaleX: scaleFactor, y: scaleFactor)
//    }
//
//    // Reset the activity indicator animation
//    func prepareInitialAnimation() {
//        isAnimatingFinal = false
//        loadMorePostsActivityIndicator.stopAnimating()
//        loadMorePostsActivityIndicator.transform = CGAffineTransform(scaleX: 0, y: 0)
//    }
//
//    func startAnimate() {
//        isAnimatingFinal = true
//        loadMorePostsActivityIndicator.startAnimating()
//    }
//
//    func stopAnimating() {
//        isAnimatingFinal = false
//        loadMorePostsActivityIndicator.stopAnimating()
//    }
//
//    func animateFinal() {
//        if isAnimatingFinal {
//            return
//        }
//        isAnimatingFinal = true
//        UIView.animate(withDuration: 1) {
//            self.loadMorePostsActivityIndicator.transform = CGAffineTransform.identity
//        }
//    }
}
