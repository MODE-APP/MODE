//
//  PostCollectionViewCell.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/15/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class PostCollectionViewCell: UICollectionViewCell {
    
    // MARK: - Outlets
    
    @IBOutlet weak var postImageImageView: UIImageView!
    
    // MARK: - Properties
    
    var image: UIImage? {
        didSet {
            updatePostImage()
        }
    }
    
    // MARK: - Custom Function
    
    func updatePostImage() {
        guard let image = image else {return}
        postImageImageView.image = image
    }
}
