//
//  notificationTableViewCell.swift
//  Mode
//
//  Created by Jackson Tubbs on 9/17/19.
//  Copyright © 2019 Jax Tubbs. All rights reserved.
//

import UIKit

class notificationTableViewCell: UITableViewCell {

    // MARK: - Outlets
    
    @IBOutlet weak var notificationSenderProfilePhoto: UIImageView!
    @IBOutlet weak var notificationPhoto: UIImageView!
    @IBOutlet weak var notificationLabel: UILabel!
    
    // MARK: - Properties
    
    
    
    // MARK: - Lifecycle
    
    override func awakeFromNib() {
        super.awakeFromNib()
        notificationSenderProfilePhoto.image = UIImage(named: "exampleProfilePhoto.jpg")
        notificationPhoto.image = UIImage(named: "profileOne.png")
        notificationLabel.text = "@telatubby Liked your photo"
    }
}
